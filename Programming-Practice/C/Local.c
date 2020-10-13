#include <stdio.h>
#include <stdlib.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <netdb.h>
#include <ifaddrs.h>
#include <string.h>
#include <malloc.h>
#include <errno.h>
#include <string.h>
#include <stdint.h>
#include<time.h>

#define BUF_SIZE 4000
#define LOCAL_CACHE "Local.txt"
#define LOCAL_TADDR "127.2.2.2"
#define LOCAL_UADDR "127.2.2.2"
#define REMOTE_UADDR "127.3.3.3"


static const uint32_t QR_MASK = 0x8000;
static const uint32_t OPCODE_MASK = 0x7800;
static const uint32_t AA_MASK = 0x0400;
static const uint32_t TC_MASK = 0x0200;
static const uint32_t RD_MASK = 0x0100;
static const uint32_t RA_MASK = 0x8000;
static const uint32_t RCODE_MASK = 0x000F;

/* Response Type */
enum {
  Ok_ResponseType = 0,
  FormatError_ResponseType = 1,
  ServerFailure_ResponseType = 2,
  NameError_ResponseType = 3,
  NotImplemented_ResponseType = 4,
  Refused_ResponseType = 5
};

/* Resource Record Types */
enum {
  A_Resource_RecordType = 1,
  NS_Resource_RecordType = 2,
  CNAME_Resource_RecordType = 5,
  SOA_Resource_RecordType = 6,
  PTR_Resource_RecordType = 12,
  MX_Resource_RecordType = 15,
  TXT_Resource_RecordType = 16,
  AAAA_Resource_RecordType = 28,
  SRV_Resource_RecordType = 33
};

/* Operation Code */
enum {
  QUERY_OperationCode = 0, /* standard query */
  IQUERY_OperationCode = 1, /* inverse query */
  STATUS_OperationCode = 2, /* server status request */
  NOTIFY_OperationCode = 4, /* request zone transfer */
  UPDATE_OperationCode = 5 /* change resource records */
};

/* Response Code */
enum {
  NoError_ResponseCode = 0,
  FormatError_ResponseCode = 1,
  ServerFailure_ResponseCode = 2,
  NameError_ResponseCode = 3
};

/* Query Type */
enum {
  IXFR_QueryType = 251,
  AXFR_QueryType = 252,
  MAILB_QueryType = 253,
  MAILA_QueryType = 254,
  STAR_QueryType = 255
};

/*
* Types.
*/

/* Question Section */
struct Question {
  char *qName;
  uint16_t qType;
  uint16_t qClass;
  struct Question* next; // for linked list
};

/* Data part of a Resource Record */
union ResourceData {
  struct {
    char *txt_data;
  } txt_record;
  struct {
    uint8_t addr[4];
  } a_record;
  struct {
    char* MName;
    char* RName;
    uint32_t serial;
    uint32_t refresh;
    uint32_t retry;
    uint32_t expire;
    uint32_t minimum;
  } soa_record;
  struct {
    char*name;
  } name_server_record;
  struct {
    char* name;
  } cname_record;
  struct {
    char *name;
  } ptr_record;
  struct {
    uint16_t preference;
    char *exchange;
  } mx_record;
  struct {
    uint8_t addr[16];
  } aaaa_record;
  struct {
    uint16_t priority;
    uint16_t weight;
    uint16_t port;
    char *target;
  } srv_record;
};

/* Resource Record Section */
struct ResourceRecord {
  char *name;
  uint16_t type;
  uint16_t class;
  uint32_t ttl;
  uint16_t rd_length;
  union ResourceData rd_data;
  struct ResourceRecord* next; // for linked list
};

struct Message { 
  //used for tcp encode&decode 
  uint16_t msg_length;
  uint16_t id; /* Identifier */

  /* Flags */
  uint16_t qr; /* Query/Response Flag */
  uint16_t opcode; /* Operation Code */
  uint16_t aa; /* Authoritative Answer Flag */
  uint16_t tc; /* Truncation Flag */
  uint16_t rd; /* Recursion Desired */
  uint16_t ra; /* Recursion Available */
  uint16_t rcode; /* Response Code */

  uint16_t qdCount; /* Question Count */
  uint16_t anCount; /* Answer Record Count */
  uint16_t nsCount; /* Authority Record Count */
  uint16_t arCount; /* Additional Record Count */

  /* At least one question; questions are copied to the response 1:1 */
  struct Question* questions;

  /*
  * Resource records to be send back.
  * Every resource record can be in any of the following places.
  * But every place has a different semantic.
  */
  struct ResourceRecord* answers;
  struct ResourceRecord* authorities;
  struct ResourceRecord* additionals;
};
//
//By Lio
//structure used for lookup dn in txt
struct RR{
        unsigned char Domain_name[45];
        unsigned int TTL;
        unsigned char Class[6];
        unsigned char type[6];
        unsigned char value[31];
};

int DomainProcess(char* DomainName,char** DivideResult){
	char* Domain;
  int len=0;
	int i=1;
	DivideResult[0]=strtok(DomainName,".");
	while(Domain!=NULL){
		Domain=strtok(NULL,".");
		if(Domain!=NULL){
			DivideResult[i]=Domain;
      len+=strlen(Domain);
			i++;
		}
	}
  return len+i+1;
}
int Domain_Count(char* Domainname){
        char* DomainName=(char*)malloc(strlen(Domainname));
        strcpy(DomainName,Domainname);
        char* Domain;
         int len=0;
        int i=1;
        Domain=strtok(DomainName,".");
        len+=strlen(Domain);
        while(Domain!=NULL){
                Domain=strtok(NULL,".");
                if(Domain!=NULL){
                        //DivideResult[i]=Domain;
                 len+=strlen(Domain);
                 i++;
                }
        }
  return len+i+1;
}

int CheckRecord(char* ip ,struct RR* re){
        FILE* fp;
        fp=fopen("cache.txt","r");
        while(!feof(fp)){
        int i=fscanf(fp,"%s %d %s %s %s",re->Domain_name,&re->TTL,re->Class,re->type,re->value);
                if(strcmp(re->Domain_name,ip)==0)
                        {
                        fclose(fp);
                        return 1;
                        }
        }
        fclose(fp);
        return 0;
}


int get_AAAA_Record(uint8_t addr[16], const char domain_name[])
{
  if (strcmp("foo.bar.com", domain_name) == 0)
  {
    addr[0] = 0xfe;
    addr[1] = 0x80;
    addr[2] = 0x00;
    addr[3] = 0x00;
    addr[4] = 0x00;
    addr[5] = 0x00;
    addr[6] = 0x00;
    addr[7] = 0x00;
    addr[8] = 0x00;
    addr[9] = 0x00;
    addr[10] = 0x00;
    addr[11] = 0x00;
    addr[12] = 0x00;
    addr[13] = 0x00;
    addr[14] = 0x00;
    addr[15] = 0x01;
    return 0;
  }
  else
  {
    return -1;
  }
}
//if we want to resolve the PTR record we have to changge the way we get the rrdata
//the chinese domain name have different type of storage,we have to get the 
struct ResourceRecord* GetRecord(char* ip ,int type,int* count){
	FILE* fp;
	struct RR* re=(struct RR*)malloc(sizeof(struct RR));
	struct ResourceRecord *rd=NULL;
  //=(struct ResourceRecord*)malloc(sizeof(struct ResourceRecord));
  //struct ResourceRecord rc*;
	fp=fopen(LOCAL_CACHE,"r");
	char* name=strdup(ip);
	printf("reader open ,name has copy %s",name);
  int caflag=0;
  int seekres=0;
  while(!seekres){
	  switch(type){
	    case A_Resource_RecordType:{
        printf("\nchecking A record\n");
	      while(1){
	        int i=fscanf(fp,"%s %d %s %s %s",re->Domain_name,&re->TTL,re->Class,re->type,re->value);
	        printf("\n%s %d %s %s %s",re->Domain_name,re->TTL,re->Class,re->type,re->value);
          if(feof(fp))break;
          if(strcmp(re->Domain_name,name)==0){
            if(strcmp(re->type,"A")==0){
              struct ResourceRecord *rr=(struct ResourceRecord*)malloc(sizeof(struct ResourceRecord));
              rr->name=strdup(re->Domain_name);
		          rr->class=1;
		          rr->type=1;
		          rr->ttl=re->TTL;
		          rr->rd_length=4;
		          char* check[4];
              int k;
              for(k=0;k<4;k++)
              check[k]=(char*)malloc(sizeof(char)*4);
		          DomainProcess(re->value,check);
              printf("%s.%s.%s.%s",check[0],check[1],check[2],check[3]);
		          rr->rd_data.a_record.addr[0]=atoi(check[0]);
		          rr->rd_data.a_record.addr[1]=atoi(check[1]);
		          rr->rd_data.a_record.addr[2]=atoi(check[2]);
		          rr->rd_data.a_record.addr[3]=atoi(check[3]);
              rr->next=NULL;
              *count+=1;
              caflag=1;
              if(rd){
                printf("there is another record before");
                rd->next=rr;
              }else{
                printf("there is no record before");
		            rd=rr;
              }
              seekres=1;
	          }
           else if(strcmp(re->type,"CNAME")==0){
              struct ResourceRecord *rr=(struct ResourceRecord*)malloc(sizeof(struct ResourceRecord));
              rr->name=strdup(re->Domain_name);
              rr->class=1;
              rr->type=5;
              rr->ttl=re->TTL;
              rr->rd_length=Domain_Count(re->value);
              rr->rd_data.cname_record.name=strdup(re->value);
              *count+=1;
              name=rr->rd_data.cname_record.name;
              type=A_Resource_RecordType;
              fseek(fp,0,SEEK_SET); 
              rd=rr;
              break; 	
	          }
            
          }
	      //printf("\n lock A record");
        };

	      if(!rd) {
	      seekres=1;
	      printf("\nCannot find record A\n");
	      }       
	      break;
        }
	    case CNAME_Resource_RecordType:
        printf("\nCheck CNAME Record\n");
        while(!feof(fp)){
	        int i=fscanf(fp,"%s %d %s %s %s",re->Domain_name,&re->TTL,re->Class,re->type,re->value);
	        if(strcmp(re->Domain_name,name)==0){
            if(strcmp(re->type,"CNAME")==0){
            struct ResourceRecord *rr=(struct ResourceRecord*)malloc(sizeof(struct ResourceRecord));
            rr->name=strdup(re->Domain_name);
            rr->class=1;
            rr->type=5;
            rr->ttl=re->TTL;
            rr->rd_length=Domain_Count(re->value);
            rr->rd_data.cname_record.name=strdup(re->value);
            name=rr->rd_data.cname_record.name;
            type=1;
            rd=rr;
            *count+=1;
            fseek(fp,0,SEEK_SET);  	
	        }
          }
        }
        if(!rd){
          seekres=1;
          printf("\nno such a record\n");
        }
    break;
	    case MX_Resource_RecordType:
      printf("\nCheck MX Record\n");
        while(!feof(fp)){
	        int i=fscanf(fp,"%s %d %s %s %s",re->Domain_name,&re->TTL,re->Class,re->type,re->value);
	        if(strcmp(re->Domain_name,name)==0){
            if(strcmp(re->type,"MX")==0){
              struct ResourceRecord *rr=(struct ResourceRecord*)malloc(sizeof(struct ResourceRecord));
              rr->name=strdup(re->Domain_name);
              rr->class=1;
              rr->type=15;
              rr->ttl=re->TTL;
              rr->rd_length=Domain_Count(re->value)+2;
              rr->rd_data.mx_record.preference=5;
              rr->rd_data.mx_record.exchange=strdup(re->value);
              name=rr->rd_data.mx_record.exchange;
              rd=rr;
              *count+=1;
              type=A_Resource_RecordType;
               
          //seekres=0;
            }
            
          }
        }
        fseek(fp,0,SEEK_SET);
        if(!rd){
          seekres=1;
          printf("\ncannnot find any record\n");
        }  
	    break;
	  case PTR_Resource_RecordType:
    printf("\n Check PTR Record\n");
    while(!feof(fp)){
	        int i=fscanf(fp,"%s %d %s %s %s",re->Domain_name,&re->TTL,re->Class,re->type,re->value);
	        if(strcmp(re->Domain_name,name)==0){
            if(strcmp(re->type,"PTR")==0){
            struct ResourceRecord *rr=(struct ResourceRecord*)malloc(sizeof(struct ResourceRecord));
            rr->name=strdup(re->Domain_name);
            rr->class=1;
            rr->type=12;
            rr->ttl=re->TTL;
            rr->rd_length=Domain_Count(re->value);
            rr->rd_data.ptr_record.name=strdup(re->value);
            rr->next=NULL;
            *count+=1;
      //rr->rd_data.mx_record.exchange=strdup(re->value);
            rd=rr;
            seekres=0;
            break;
            }
            
          }
        }
    seekres=1; 
	  break;
     default:
    printf("\nget Record error of cannot find this record\n");
    seekres=1;
    }
   
  }
  fclose(fp);
  return rd;

}
/*
* Debugging functions.
*/

void print_hex(uint8_t* buf, size_t len)
{
  int i;
  printf("%zu bytes:\n", len);
  for(i = 0; i < len; ++i)
    printf("%02x ", buf[i]);
  printf("\n"); 
}

void print_resource_record(struct ResourceRecord* rr)
{
  int i;
  while (rr)
  {
    printf("  ResourceRecord { name '%s', type %u, class %u, ttl %u, rd_length %u, ",
        rr->name,
        rr->type,
        rr->class,
        rr->ttl,
        rr->rd_length
   );

    union ResourceData *rd = &rr->rd_data;
    switch (rr->type)
    {
      case A_Resource_RecordType:
        printf("Address Resource Record { address ");

        for(i = 0; i < 4; ++i)
          printf("%s%u", (i ? "." : ""), rd->a_record.addr[i]);

        printf(" }");
        break;
      case NS_Resource_RecordType:
        printf("Name Server Resource Record { name %s }",
          rd->name_server_record.name
       );
        break;
      case CNAME_Resource_RecordType:
        printf("Canonical Name Resource Record { name %u }",
          rd->cname_record.name
       );
        break;
      case SOA_Resource_RecordType:
        printf("SOA { MName '%s', RName '%s', serial %u, refresh %u, retry %u, expire %u, minimum %u }",
          rd->soa_record.MName,
          rd->soa_record.RName,
          rd->soa_record.serial,
          rd->soa_record.refresh,
          rd->soa_record.retry,
          rd->soa_record.expire,
          rd->soa_record.minimum
       );
        break;
      case PTR_Resource_RecordType:
        printf("Pointer Resource Record { name '%s' }",
          rd->ptr_record.name
       );
        break;
      case MX_Resource_RecordType:
        printf("Mail Exchange Record { preference %u, exchange '%s' }",
          rd->mx_record.preference,
          rd->mx_record.exchange
       );
        break;
      case TXT_Resource_RecordType:
        printf("Text Resource Record { txt_data '%s' }",
          rd->txt_record.txt_data
       );
        break;
      case AAAA_Resource_RecordType:
        printf("AAAA Resource Record { address ");

        for(i = 0; i < 16; ++i)
          printf("%s%02x", (i ? ":" : ""), rd->aaaa_record.addr[i]);

        printf(" }");
        break;
      default:
        printf("Unknown Resource Record { ??? }");
    }
    printf("}\n");
    rr = rr->next;
  }
}

void print_query(struct Message* msg)
{
  printf("QUERY { ID: %02x", msg->id);
  printf(". FIELDS: [ QR: %u, OpCode: %u ]", msg->qr, msg->opcode);
  printf(", QDcount: %u", msg->qdCount);
  printf(", ANcount: %u", msg->anCount);
  printf(", NScount: %u", msg->nsCount);
  printf(", ARcount: %u,\n", msg->arCount);

  struct Question* q = msg->questions;
  while (q)
  {
    printf("  Question { qName '%s', qType %u, qClass %u }\n",
      q->qName,
      q->qType,
      q->qClass
   );
    q = q->next;
  }

  print_resource_record(msg->answers);
  print_resource_record(msg->authorities);
  print_resource_record(msg->additionals);

  printf("}\n");
}


/*
* Basic memory operations.
*/

uint16_t get16bits(const uint8_t** buffer)
{
  uint16_t value;

  memcpy(&value, *buffer, 2);
  *buffer += 2;

  return ntohs(value);
}
uint8_t get8bits(const uint8_t** buffer){
  uint8_t value;
  memcpy(&value,*buffer,1);
  *buffer+=1;
  return value;
}
uint32_t get32bits(const uint8_t** buffer){
  uint32_t value;
  memcpy(&value,*buffer,4);
  *buffer+=4;
  return ntohl(value);
}
void put8bits(uint8_t** buffer, uint8_t value)
{
  memcpy(*buffer, &value, 1);
  *buffer += 1;
}

void put16bits(uint8_t** buffer, uint16_t value)
{
  value = htons(value);
  memcpy(*buffer, &value, 2);
  *buffer += 2;
}

void put32bits(uint8_t** buffer, uint32_t value)
{
  value = htons(value);
  memcpy(*buffer, &value, 4);
  *buffer += 4;
}


/*
* Deconding/Encoding functions.
*/

// 3foo3bar3com0 => foo.bar.com
char* decode_domain_name(const uint8_t** buffer)
{
  char name[256];
  const uint8_t* buf = *buffer;
  int j = 0;
  int i = 0;

  while (buf[i] != 0)
  {
    //if (i >= buflen || i > sizeof(name))
    //  return NULL;

    if (i != 0)
    {
      name[j] = '.';
      j += 1;
    }

    int len = buf[i];
    i += 1;

    memcpy(name+j, buf+i, len);
    i += len;
    j += len;
  }

  name[j] = '\0';

  *buffer += i + 1; //also jump over the last 0

  return strdup(name);
}

// foo.bar.com => 3foo3bar3com0
void encode_domain_name(uint8_t** buffer, const char* domain)
{
  uint8_t* buf = *buffer;
  const char* beg = domain;
  const char* pos;
  int len = 0;
  int i = 0;

  while ((pos = strchr(beg, '.')))
  {
    len = pos - beg;
    buf[i] = len;
    i += 1;
    memcpy(buf+i, beg, len);
    i += len;

    beg = pos + 1;
  }

  len = strlen(domain) - (beg - domain);

  buf[i] = len;
  i += 1;

  memcpy(buf + i, beg, len);
  i += len;

  buf[i] = 0;
  i += 1;

  *buffer += i;
}


void decode_header(struct Message* msg, const uint8_t** buffer)
{
  msg->msg_length=get16bits(buffer);
  msg->id = get16bits(buffer);

  uint32_t fields = get16bits(buffer);
  msg->qr = (fields & QR_MASK) >> 15;
  msg->opcode = (fields & OPCODE_MASK) >> 11;
  msg->aa = (fields & AA_MASK) >> 10;
  msg->tc = (fields & TC_MASK) >> 9;
  msg->rd = (fields & RD_MASK) >> 8;
  msg->ra = (fields & RA_MASK) >> 7;
  msg->rcode = (fields & RCODE_MASK) >> 0;

  msg->qdCount = get16bits(buffer);
  msg->anCount = get16bits(buffer);
  msg->nsCount = get16bits(buffer);
  msg->arCount = get16bits(buffer);
}
//caculate rr field
uint16_t ResourceField_Caculate(struct Message* msg){
  uint16_t len=0;
  printf("\ncaculating rr field \n");
  if(msg->anCount!=0)
     len+=(uint16_t)count_resource_records(msg->answers);
  if(msg->nsCount!=0)
     len+=(uint16_t)count_resource_records(msg->authorities);
  if(msg->arCount!=0)
     len+=(uint16_t)count_resource_records(msg->additionals);
  printf("\ncaculate finish\n");
  return len;
}

void encode_header(struct Message* msg, uint8_t** buffer)
{
  //msg->msg_length=(uint16_t)(12+msg->qdCount*8+
  //(msg->qdCount+msg->anCount+msg->nsCount+msg->arCount)*16);
  printf("msg length%d",msg->msg_length);
  
  //printf("!!!!msg length:%d",msg->msg_length);
  //msg->msg_length=msg->msg_length+msg->anCount*(sizeof(struct ResourceData)+sizeof(struct ResourceRecord));
  
  put16bits(buffer,msg->msg_length);
  put16bits(buffer, msg->id);

  int fields = 0;
  fields |= (msg->qr << 15) & QR_MASK;
  fields |= (msg->rcode << 0) & RCODE_MASK;
  // TODO: insert the rest of the fields
  put16bits(buffer, fields);

  put16bits(buffer, msg->qdCount);
  put16bits(buffer, msg->anCount);
  put16bits(buffer, msg->nsCount);
  put16bits(buffer, msg->arCount);
}

int encode_resource_records(struct ResourceRecord* rr, uint8_t** buffer)
{
  int i;
  //printf("\n\n encode below \n\n");
  print_resource_record(rr);
  while (rr)
  {
    //printf("\nencoding rr\n");
    // Answer questions by attaching resource sections.
    encode_domain_name(buffer, rr->name);
    put16bits(buffer, rr->type);
    put16bits(buffer, rr->class);
    put32bits(buffer, rr->ttl);
    put16bits(buffer, rr->rd_length);
    
    switch (rr->type)
    {
      case A_Resource_RecordType:
        for(i = 0; i < 4; ++i)
          put8bits(buffer, rr->rd_data.a_record.addr[i]);
        break;
      case AAAA_Resource_RecordType:
        for(i = 0; i < 16; ++i)
          put8bits(buffer, rr->rd_data.aaaa_record.addr[i]);
        break;
      case CNAME_Resource_RecordType:
        //put_rdstring_in_buffer(buffer,rr->rd_data.cname_record.name);
        encode_domain_name(buffer,rr->rd_data.cname_record.name);
        break;
      case MX_Resource_RecordType:
          put16bits(buffer,rr->rd_data.mx_record.preference);
          //put_rdstring_in_buffer(buffer,rr->rd_data.mx_record.exchange);
          encode_domain_name(buffer,rr->rd_data.mx_record.exchange);
          break;
      case PTR_Resource_RecordType:
          //put_rdstring_in_buffer(buffer,rr->rd_data.ptr_record.name);
          encode_domain_name(buffer,rr->rd_data.ptr_record.name);
      default:
        fprintf(stderr, "Unknown type %u. => Ignore resource record.\n", rr->type);
      return 1;
    }
    //printf("\n put in success\n");
    rr = rr->next;
  }

  return 0;
}
/* @return 0 upon failure, 1 upon success */
int encode_msg(struct Message* msg, uint8_t** buffer)
{
  //printf("\necoding msg\n");
  struct Question* q;
  int rc;

  encode_header(msg, buffer);
  //printf("\nencod header finish\n");
  q = msg->questions;
  while(q)
  {
    encode_domain_name(buffer, q->qName);
    put16bits(buffer, q->qType);
    put16bits(buffer, q->qClass);

    q = q->next;
  }
  //printf("\nencod question finish\n");
  rc = 0;
  rc |= encode_resource_records(msg->answers, buffer);
  //printf("\nresult of encoding answers:%d\n",rc);
  rc |= encode_resource_records(msg->authorities, buffer);
  rc |= encode_resource_records(msg->additionals, buffer);

  return rc;
}

int decode_msg(struct Message* msg, const uint8_t* buffer, int size)
{
  int i;

  decode_header(msg, &buffer);

  if (msg->anCount != 0 || msg->nsCount != 0)
  {
    printf("Only questions expected!\n");
    return -1;
  }

  // parse questions
  uint32_t qcount = msg->qdCount;
  struct Question* qs = msg->questions;
  for (i = 0; i < qcount; ++i)
  {
    struct Question* q = malloc(sizeof(struct Question));

    q->qName = decode_domain_name(&buffer);
    q->qType = get16bits(&buffer);
    q->qClass = get16bits(&buffer);

    // prepend question to questions list
    q->next = qs;
    msg->questions = q;
  }

  // We do not expect any resource records to parse here.

  return 0;
}

// For every question in the message add a appropiate resource record
// in either section 'answers', 'authorities' or 'additionals'.
int resolver_process(struct Message* msg)
{
  int resolve_result=0;
  struct ResourceRecord* beg;
  struct ResourceRecord* rr;
  struct Question* q;
  int rc;

  // leave most values intact for response
  msg->qr = 1; // this is a response
  msg->aa = 1; // this server is authoritative
  msg->ra = 0; // no recursion available
  msg->rcode = Ok_ResponseType;

  // should already be 0
  msg->anCount = 0;
  msg->nsCount = 0;
  msg->arCount = 0;

  // for every question append resource records
  q = msg->questions;
  while (q)
  {
    //rr = malloc(sizeof(struct ResourceRecord));
    //memset(rr, 0, sizeof(struct ResourceRecord));
    
    //rr->name=strdup(q->qName);
    //printf("\nrr name:%s \nqname:%s",rr->name,q->qName);
    //rr->name = strdup(q->qName);
    //rr->type = q->qType;
    //rr->class = q->qClass;
    //rr->ttl = 60*60; // in seconds; 0 means no caching

    printf("Query for '%s'\n", q->qName);

    // We only can only answer two question types so far
    // and the answer (resource records) will be all put
    // into the answers list.
    // This behavior is probably non-standard!
    
    switch (q->qType)
    {
      case A_Resource_RecordType:
        {int count=0;
        rr=GetRecord(q->qName,q->qType,&count);
        printf(" \ncheck result !\n");
        print_resource_record(rr);
        if (count==0)
        {
          //free(rr->name);
          //free(rr);
          printf("cannot find %s",q->qName);
          goto next;
        }
        else{
          msg->anCount=count;
          resolve_result+=count;
        }
        
        break;
        }
      case MX_Resource_RecordType:
        {int count=0;
        rr=GetRecord(q->qName,q->qType,&count);
        if (count==0)
        {
          //free(rr->name);
          //free(rr);
          printf("cannot find %s",q->qName);
          goto next;
        }
        else{
          msg->anCount=count-1;
          msg->arCount=1;
          msg->additionals=rr->next;
          
          rr->next=NULL;
          
          resolve_result+=count;
        }
        break;}
      case PTR_Resource_RecordType:
        {
        int count=0;
        rr=GetRecord(q->qName,q->qType,&count);
        if (count==0)
        {
          //free(rr->name);
          //free(rr);
          printf("cannot find %s",q->qName);
          goto next;
        }
        else{
          msg->anCount=1;
          resolve_result+=count;
        }
        break;
        }
      /*
      case NS_Resource_RecordType:
      case CNAME_Resource_RecordType:
      case SOA_Resource_RecordType:
      case PTR_Resource_RecordType:
      case MX_Resource_RecordType:
      case TXT_Resource_RecordType:
      */
      default:
        free(rr);
        msg->rcode = NotImplemented_ResponseType;
        printf("Cannot answer question of type %d.\n", q->qType);
        goto next;
    }

    //msg->anCount++;

    // prepend resource record to answers list
    //beg = msg->answers;
    msg->answers = rr;
    //rr->next = beg;

    // jump here to omit question,I can't find another way better,so I use goto .
    next:

    // process next question
    q = q->next;
  }
  //printf("\nresolve print\n");
  //print_resource_record(msg->answers);
  //print_resource_record(msg->additionals);
  msg->msg_length=msg->msg_length+ResourceField_Caculate(msg);
  
  return resolve_result;
}

/* @return 0 upon failure, n upon success */

int count_resource_records(struct ResourceRecord* rr){
  int len=0;
  while(rr){
    len+=Domain_Count(rr->name);
    if(rr->type)
      len+=2;
    if(rr->class)
      len+=2;
    if(rr->ttl)
      len+=4;
    if(rr->rd_length)
      len+=2;
    switch(rr->type)
    {
      case A_Resource_RecordType:
      len+=4;
      break;
      case CNAME_Resource_RecordType:
      len=len+Domain_Count(rr->rd_data.cname_record.name);
      //CNAME record have to count the byte 
      break;
      case MX_Resource_RecordType:
      len=len+Domain_Count(rr->rd_data.mx_record.exchange)+2;
      break;
      case PTR_Resource_RecordType:
      len+=Domain_Count(rr->rd_data.ptr_record.name);
      break;
      default:
      printf("cannot find kind\n");
    }
    rr=rr->next;  
  }
  printf("\nRecord Count %d",len);
  return len;
}

void free_questions(struct Question* qq){
  struct Question* next;
  while(qq){
    free(qq->qName);
    next=qq->next;
    free(qq);
    qq=next;
  }
}
void free_resource_records(struct ResourceRecord* rr)
{
  struct ResourceRecord* next;

  while(rr){
    free(rr->name);
    next=rr->next;
    free(rr);
    rr=next;
  }
}

//following is the UDP part
struct udp_Message {
  uint16_t id; /* Identifier */
  /* Flags */
  uint16_t qr; /* Query/Response Flag */
  uint16_t opcode; /* Operation Code */
  uint16_t aa; /* Authoritative Answer Flag */
  uint16_t tc; /* Truncation Flag */
  uint16_t rd; /* Recursion Desired */
  uint16_t ra; /* Recursion Available */
  uint16_t rcode; /* Response Code */

  uint16_t qdCount; /* Question Count */
  uint16_t anCount; /* Answer Record Count */
  uint16_t nsCount; /* Authority Record Count */
  uint16_t arCount; /* Additional Record Count */

  /* At least one question; questions are copied to the response 1:1 */
  struct Question* questions;

  /*
  * Resource records to be send back.
  * Every resource record can be in any of the following places.
  * But every place has a different semantic.
  */
  struct ResourceRecord* answers;
  struct ResourceRecord* authorities;
  struct ResourceRecord* additionals;
};
void udp_print_query(struct udp_Message* msg)
{
  printf("QUERY ID: %02x\n", msg->id);
  printf("FIELDS: [ QR: %u, OpCode: %u ]\n", msg->qr, msg->opcode);
  printf("QDcount: %u\n", msg->qdCount);
  printf("ANcount: %u\n", msg->anCount);
  printf("NScount: %u\n", msg->nsCount);
  printf("ARcount: %u,\n", msg->arCount);

  struct Question* q = msg->questions;
  while (q)
  {
    printf("  Question { qName '%s', qType %u, qClass %u }\n",
      q->qName,
      q->qType,
      q->qClass
   );
    q = q->next;
  }

  print_resource_record(msg->answers);
  print_resource_record(msg->authorities);
  print_resource_record(msg->additionals);

  printf("}\n");
}
void udp_decode_header(struct udp_Message* msg, const uint8_t** buffer)
{
  msg->id = get16bits(buffer);

  uint32_t fields = get16bits(buffer);
  msg->qr = (fields & QR_MASK) >> 15;
  msg->opcode = (fields & OPCODE_MASK) >> 11;
  msg->aa = (fields & AA_MASK) >> 10;
  msg->tc = (fields & TC_MASK) >> 9;
  msg->rd = (fields & RD_MASK) >> 8;
  msg->ra = (fields & RA_MASK) >> 7;
  msg->rcode = (fields & RCODE_MASK) >> 0;

  msg->qdCount = get16bits(buffer);
  msg->anCount = get16bits(buffer);
  msg->nsCount = get16bits(buffer);
  msg->arCount = get16bits(buffer);
}
void udp_encode_header(struct udp_Message* msg, uint8_t** buffer)
{
  put16bits(buffer, msg->id);

  int fields = 0;
  fields |= (msg->qr << 15) & QR_MASK;
  fields |= (msg->rcode << 0) & RCODE_MASK;
  // TODO: insert the rest of the fields
  put16bits(buffer, fields);

  put16bits(buffer, msg->qdCount);
  put16bits(buffer, msg->anCount);
  put16bits(buffer, msg->nsCount);
  put16bits(buffer, msg->arCount);
}

int udp_encode_msg(struct udp_Message* msg, uint8_t** buffer)
{
  struct Question* q;
  int rc;

  udp_encode_header(msg, buffer);

  q = msg->questions;
    while (q)
    {
    encode_domain_name(buffer, q->qName);
    put16bits(buffer, q->qType);
    put16bits(buffer, q->qClass);

    q = q->next;
    }

  rc = 0;
  rc |= encode_resource_records(msg->answers, buffer);
  rc |= encode_resource_records(msg->authorities, buffer);
  rc |= encode_resource_records(msg->additionals, buffer);

  return rc;
}

int decode_resource_records(struct ResourceRecord** rr, const uint8_t** buffer,int Count)
{
  if(Count==0)return 0;
  printf("\ndecoding resource records\n");
  int i,j;
  int rrcount=Count;
  struct ResourceRecord* as=*rr;
  printf("\n decode record \n");
  for(i=0;i<rrcount;i++)
  {
    struct ResourceRecord* r=(struct ResourceRecord*)malloc(sizeof(struct ResourceRecord));
    r->name=decode_domain_name(buffer);
    r->type=get16bits(buffer);
    r->class=get16bits(buffer);
    r->ttl=get32bits(buffer);
    r->rd_length=get16bits(buffer);
    switch(r->type){
    case A_Resource_RecordType:
    for(j=0;j<4;j++)
      r->rd_data.a_record.addr[j]=get8bits(buffer);
    break;
    case CNAME_Resource_RecordType:
      r->rd_data.cname_record.name=decode_domain_name(buffer);
      break;
    case MX_Resource_RecordType:
    r->rd_data.mx_record.preference=get16bits(buffer);
    r->rd_data.mx_record.exchange=decode_domain_name(buffer);
    break;
    case PTR_Resource_RecordType:
      r->rd_data.ptr_record.name=decode_domain_name(buffer);
    break;
    //case MX_Resource_RecordType:
    }
    r->next=as;
    *rr=r;
    as=*rr;
    printf("decode success");
  }
  //*rr=as;
  print_resource_record(*rr);
  return 1;
}
  
int udp_decode_msg(struct udp_Message* msg, const uint8_t* buffer, int size)
{ 
  int i;
  int j;
  udp_decode_header(msg, &buffer);
  size=size-12;

  // parse questions
  uint32_t qcount = msg->qdCount;
  struct Question* qs = msg->questions;
  for (i = 0; i < qcount; ++i)
  {
    struct Question* q = malloc(sizeof(struct Question));
    q->qName = decode_domain_name(&buffer);
    q->qType = get16bits(&buffer);
    q->qClass = get16bits(&buffer);
    size=size-strlen(q->qName)-1-2-2;
    // prepend question to questions list
    q->next = qs;
    msg->questions = q;
  }
  printf("start");
  msg->answers=NULL;
  msg->additionals=NULL;
  msg->authorities=NULL;
  printf("\n answers:\n");
  decode_resource_records(&msg->answers,&buffer,msg->anCount);
  print_resource_record(msg->answers);
  decode_resource_records(&msg->authorities,&buffer,msg->nsCount);
  printf("\n authorities:\n");
  print_resource_record(msg->authorities);
  decode_resource_records(&msg->additionals,&buffer,msg->arCount);
  printf("\n additionals:\n");
  print_resource_record(msg->additionals);
    printf("decode success");
    return 0;
}

int ResourceRecord_copy(struct ResourceRecord* rr,struct ResourceRecord* rd){
 printf("\nresource record copy\n");
 rr->name=strdup(rd->name);
 rr->type=rd->type;
 rr->ttl=rd->ttl;
 rr->class=rd->class;
 rr->rd_length=rd->rd_length;
 rr->rd_data=rd->rd_data;
}

int Message_copy(struct Message* tcp_msg,struct udp_Message* udp_msg){
  udp_msg->id=tcp_msg->id;
  udp_msg->qr=0;
  udp_msg->opcode=tcp_msg->opcode;
  udp_msg->aa=tcp_msg->aa;
  udp_msg->tc=tcp_msg->tc;
  udp_msg->rd=0;//local dns to servers are not allow to recursion
  udp_msg->ra=tcp_msg->ra;
  udp_msg->rcode=tcp_msg->rcode;
  udp_msg->qdCount=tcp_msg->qdCount;
  udp_msg->anCount=tcp_msg->anCount;
  udp_msg->nsCount=0;
  udp_msg->arCount=tcp_msg->arCount;
  udp_msg->questions=tcp_msg->questions;

}

//for local server get ip from local or Recorddata
//the 2nd argument must malloc enough storage,at least 16 byte
//the result can be used for sock_addr.  =inet_addr(result)
void getIPfromRecord(unsigned char addr[4],char** result){
        unsigned char* name;
        name=(char*)malloc(sizeof(char)*18);
        strcpy(name,"");
        int i=0;
        char* a[4];
        //int j=0;
        for(i;i<4;i++){
        //if(i==0)
        //a[i]=(char*)malloc(sizeof(char)*5);
        //relse
        a[i]=(char*)malloc(sizeof(char)*5);
        sprintf(a[i],"%d",addr[i]);
        printf("input:%s ",a[i]);
        }
        i=0;
        for(i;i<4;i++){
                strcat(name,a[i]);
                if(i<3)
                strcat(name,".");
        }
        //printf("name:%s",name);
        //strcpy(result,name);
        *result=strdup(name);
        printf("\nresult=%s\n",*result);
}


//udp process for localserver to server
int udp_main_process(struct Message* tcp_message){
  printf("\nudp start\n");
  // buffer for input/output binary packet
  uint8_t buffer[BUF_SIZE];
  //backup Question for recover the question
  struct Question* backup_query;
  backup_query=(struct Question*)malloc(sizeof(struct Question*));
  backup_query=tcp_message->questions;
  
  socklen_t addr_len = sizeof(struct sockaddr_in);
  struct sockaddr_in addr;
  int nbytes, rc;
  int sock;
  int port = 53;
  struct udp_Message msg;
  memset(&msg, 0, sizeof(struct udp_Message));
  
  struct Message tcp_copy;
  memset(&tcp_copy,0,sizeof(struct Message));
  tcp_copy=*tcp_message;

  struct udp_Message query;
   memset(&query, 0, sizeof(struct udp_Message));
  Message_copy(tcp_message,&query);
  Message_copy(tcp_message,&msg);
  //local addr
  addr.sin_family = AF_INET;
  addr.sin_addr.s_addr = inet_addr(LOCAL_UADDR);
  addr.sin_port = htons(port);
  //server addr
  struct sockaddr_in server_addr;
  server_addr.sin_family=AF_INET;
  server_addr.sin_addr.s_addr=inet_addr(REMOTE_UADDR);
  server_addr.sin_port=htons(port);
  sock = socket(AF_INET, SOCK_DGRAM, 0);
  printf("socket create success\n");
  int on=1;
  if((setsockopt(sock,SOL_SOCKET,SO_REUSEADDR,&on,sizeof(on)))<0){
    printf("Set sock error:%s\n",strerror(errno));
  }
  printf("Set success\n");
  rc = bind(sock, (struct sockaddr*) &addr, addr_len);
  if (rc != 0)
  {
    printf("Could not bind: %s\n", strerror(errno));
    return -1;
  }
  printf("bind success\n");
  char* server_locate=(char*)malloc(sizeof(char)*18);
  //memset(&server_locate,0,sizeof(char)*18);
  int querysuccess=0;
  //if we cannot get the answer we want ,we continue to ask another DNS for help.
  //we don't have to change the Query packet
  //if we got the right answer,stop,and feed back to the tcp;
  printf("s\n request for root 127.3.3.3\n");
  while(querysuccess!=1){
    uint8_t *l=buffer;
    udp_encode_msg(&query, &l);
    int sendlen=l-buffer;
    //udp_print_query(&query);
    clock_t start,end;

    start=clock();
    sendto(sock, buffer,sendlen, 0, (struct sockaddr*) &server_addr, addr_len);
    //free_questions(msg.questions);
    free_resource_records(msg.answers);
    free_resource_records(msg.authorities);
    //free_resource_records(msg.additionals);
    msg.questions=NULL;
    msg.answers=NULL;
    msg.authorities=NULL;
    msg.additionals=NULL;
    memset(&msg, 0, sizeof(struct udp_Message));
    printf("Listening on port %u.\n", port);

    
    nbytes = recvfrom(sock, buffer, sizeof(buffer), 0, (struct sockaddr *) &server_addr, &addr_len);
    end=clock();
    double timers=(double)(end-start)/CLOCKS_PER_SEC;
    printf("response in %f s",timers);

    printf("\nreceive %d bytes from dns server\n",nbytes);
    if (udp_decode_msg(&msg, buffer, nbytes) != 0) {
      printf("\nudp decode fail\n");
      //continue;
    }
    //printf("\nudp receive: \n");
    /* Print query */
    //udp_print_query(&msg);

    //if(udp_resolver_process(&msg)==1){
    if(msg.anCount!=0){
      switch(msg.questions->qType){
      case A_Resource_RecordType:
      tcp_message->answers=msg.answers;
      tcp_message->anCount=msg.anCount;
      msg.answers=NULL;
       tcp_message->authorities=NULL;
       tcp_message->additionals=NULL;
      tcp_message->msg_length=tcp_message->msg_length+count_resource_records(tcp_message->answers);
      break;
      case CNAME_Resource_RecordType:
      tcp_message->answers=msg.answers;
      tcp_message->anCount=msg.anCount;
      tcp_message->additionals=NULL;
      tcp_message->authorities=NULL;
      tcp_message->msg_length=tcp_message->msg_length+count_resource_records(tcp_message->answers);
      break;
      case MX_Resource_RecordType:
      tcp_message->anCount=msg.anCount;
      tcp_message->arCount=msg.arCount;
      tcp_message->answers=msg.answers;
      msg.answers=NULL;
      tcp_message->additionals=msg.additionals;
      msg.additionals=NULL;
      tcp_message->authorities=NULL;
      tcp_message->msg_length=tcp_message->msg_length+count_resource_records(tcp_message->answers)+count_resource_records(tcp_message->additionals);
      break;
      case PTR_Resource_RecordType:
      tcp_message->anCount=msg.anCount;
      tcp_message->answers=msg.answers;
      tcp_message->additionals=NULL;
      tcp_message->authorities=NULL;
      msg.answers=NULL;
      tcp_message->msg_length=tcp_message->msg_length+count_resource_records(tcp_message->answers);
      break;
      }
      return 1;
      }
    else if(msg.nsCount!=0){
      printf("\nNEXT udp process\n");
      getIPfromRecord(msg.authorities->rd_data.a_record.addr,&server_locate);
      if(server_locate==NULL){
      printf("\nserver_locate is null,cannot analyse!!\n");
      break;
      }
      printf("next addr:%s",server_locate);
      server_addr.sin_addr.s_addr=inet_addr(server_locate);
      }
    else 
      {
      printf("\nCannot find any other record\n");
      break;
      }
  }
  printf("\nUDP finish\n");
  close(sock);
  return 0;
}
int CacheRR(struct ResourceRecord* rr){
	FILE* fp;
	fp=fopen(LOCAL_CACHE,"a+");
	if(fp==NULL){
	  printf("\ncache fail\n");
    exit(1);
	  }
	else {
    while(rr){
	    switch(rr->type){
	      case A_Resource_RecordType:{
	        char* result;
          getIPfromRecord(rr->rd_data.a_record.addr,&result);
          fprintf(fp,"\n%s %d IN A %s",rr->name,rr->ttl,result);
	        free(result);
          break;
        }
        case MX_Resource_RecordType:
          fprintf(fp,"\n%s %d IN MX %s",rr->name,rr->ttl,rr->rd_data.mx_record.exchange);
          break;
        case CNAME_Resource_RecordType:
          fprintf(fp,"\n%s %d IN CNAME %s",rr->name,rr->ttl,rr->rd_data.cname_record.name);
          break;
        case PTR_Resource_RecordType:
          fprintf(fp,"\n%s %d IN PTR %s",rr->name,rr->ttl,rr->rd_data.ptr_record.name);
          break;
        }
      rr=rr->next;
    }
  
	}
  fclose(fp);
  return 1;
};
int main()
{ 
  // buffer for input/output binary packet
  uint8_t buffer[BUF_SIZE];
  struct sockaddr_in client_addr;
  socklen_t addr_len = sizeof(struct sockaddr_in);
  struct sockaddr_in addr,client;
  int nbytes, rc;
  int sock;
  int port = 53;

  struct Message msg;
  memset(&msg, 0, sizeof(struct Message));

  addr.sin_family = AF_INET;
  addr.sin_addr.s_addr = INADDR_ANY;
  addr.sin_port = htons(port);

  sock = socket(AF_INET, SOCK_STREAM, 0);
   int on=1;
  if((setsockopt(sock,SOL_SOCKET,SO_REUSEADDR,&on,sizeof(on)))<0){
    printf("Set sock error:%s\n",strerror(errno));
  }
  rc = bind(sock, (struct sockaddr*) &addr, addr_len);
  listen(sock,3);
  if (rc != 0)
  {
    printf("Could not bind: %s\n", strerror(errno));
    return 1;
  }

  printf("Listening on port %u.\n", port);
  int client_len=sizeof(client_addr);
  int conn=accept(sock,(struct sockaddr*)&client_addr,&client_len);
    
 // while (1)
  //{
    free_questions(msg.questions);
    free_resource_records(msg.answers);
    free_resource_records(msg.authorities);
    free_resource_records(msg.additionals);
    memset(&msg, 0, sizeof(struct Message));
    
   // nbytes = recvfrom(sock, buffer, sizeof(buffer), 0, (struct sockaddr *) &client_addr, &addr_len);
    nbytes=recv(conn,buffer,sizeof(buffer),0);
   // printf("%s\n %d bytes\n",buffer,nbytes);
    if(nbytes==0){
   	  printf("\n no input\n");
     }
    if (decode_msg(&msg, buffer, nbytes) != 0) {
     printf("\n cannot decode\n");
    }

    /* Print query */
    print_query(&msg);
    if(resolver_process(&msg)==0){
      udp_main_process(&msg);
      CacheRR(msg.answers);
      CacheRR(msg.authorities);
      CacheRR(msg.additionals);
      printf("\nudp finish\n");
    };

    /* Print response */
   // printf("\nThe response to Client\n");
   // print_query(&msg);
    printf("\nresponse to Client\n");
    uint8_t *p = buffer;
    if(encode_msg(&msg, &p) != 0) {
      printf("\n encode Error\n");
    }
   
   // print_query(&msg);
    printf("msg length to send %d",msg.msg_length);
     write(conn,buffer,msg.msg_length+2);
    //send(conn,buffer,sizeof(msg),0);
   // sendto(sock, buffer, buflen, 0, (struct sockaddr*) &client_addr, addr_len);
  //}
}

