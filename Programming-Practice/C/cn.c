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

#define BUF_SIZE 1500
#define LOCAL_UADDR "127.4.4.1"
#define LOCAL_CACHE "cn.txt"
/*
Root Server ,it's not different from top Servers.
*/


/*
* Masks and constants.
*/

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
    char *name;
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
  uint16_t ttl;
  uint16_t rd_length;
  union ResourceData rd_data;
  struct ResourceRecord* next; // for linked list
};

struct Message {
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
//the function used for get the domain name;
//level means the domian's level
//top must be malloc enough storage at least 14 byte

void getDomainName(char* domainname,char* top,int level){

        char* Domain;

        int len=strlen(domainname);
        Domain=(char*)malloc(len);
        strcpy(Domain,domainname);

        int i=0;
        int j=0;
        for(i;i<strlen(Domain);i++){
        if(Domain[i]=='.'){
            j++;
        }
        }
        i=0;
       for(i;i<len;i++){
                if(domainname[i]=='.')
                         j--;
                        Domain++;
                 if(j<=(level-1))
                        break;
       }
        if(level==1)
        {
        char* temp=(char*)malloc(sizeof(char)*22);
        *temp='\0';
        strcat(temp,"*.*.");
        strcat(temp,Domain);
         strcpy(top,temp);
        free(temp);
        }
       if(level==2){
        char* temp=(char*)malloc(sizeof(char)*22);
        *temp='\0';
        strcat(temp,"*.");
        strcat(temp,Domain);
        strcpy(top,temp);
        free(temp);
        }
       }

void IPProcess(char* DomainName,char** DivideResult){
  char* Domain;
	int i=1;
	DivideResult[0]=strtok(DomainName,".");
	while(Domain!=NULL){
		Domain=strtok(NULL,".");
		if(Domain!=NULL){
			DivideResult[i]=Domain;
			i++;
		}
	}
}
int CheckRecord(char* ip ,struct RR* re){
        FILE* fp;
        fp=fopen(LOCAL_CACHE,"r");
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

int get_DNS_Record(uint8_t addr[4],  char** domain_name){
  struct RR* rr;
  rr=(struct RR*)malloc(sizeof(struct RR));
  char* ns=(char*)malloc(sizeof(char)*22);
  //memset(&ns,0,sizeof(char)*22);
  //strcpy(ns,domain_name);
  getDomainName(*domain_name,ns,2);
  *domain_name=strdup(ns);
  printf("\nget domain name:%s\n",*domain_name);
  if(CheckRecord(ns,rr)==1){
  //strcpy(domain_name,ns);
  printf("%s %d %s %s %s",rr->Domain_name,rr->TTL,rr->Class,rr->type,rr->value);
  char* name[4];
  IPProcess(rr->value,name);
  addr[0]=atoi(name[0]);
  addr[1]=atoi(name[1]);
  addr[2]=atoi(name[2]);
  addr[3]=atoi(name[3]);
  free(rr);
  free(ns);
  return 0;
  }
  else {
  printf("\ncannot find SLD DNS  \n");
  free(rr);
  free(ns);
  return -1;
  }
}
int get_A_Record(uint8_t addr[4],  char* domain_name){
  struct RR* rr;
  rr=(struct RR*)malloc(sizeof(struct RR));
  if(CheckRecord(domain_name,rr)==1){
  char* name[4];
  IPProcess(rr->value,name);
  addr[0]=atoi(name[0]);
  addr[1]=atoi(name[1]);
  addr[2]=atoi(name[2]);
  addr[3]=atoi(name[3]);
  return 0;
  }
else return -1;

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

size_t get16bits(const uint8_t** buffer)
{
  uint16_t value;

  memcpy(&value, *buffer, 2);
  *buffer += 2;

  return ntohs(value);
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

void encode_header(struct Message* msg, uint8_t** buffer)
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
// For every question in the message add a appropiate resource record
// in either section 'answers', 'authorities' or 'additionals'.
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
int resolver_process(struct Message* msg)
{
  int resolve_result=0;
  struct ResourceRecord* beg;
  struct ResourceRecord* rr;
  struct Question* q;
  int rc,rdns;

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

    // We only can only answer four question types so far
    // and the answer (resource records) will be all put
    // into the answers field and additionals field.
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
          struct ResourceRecord* rd=(struct ResourceRecord*)malloc(sizeof(struct ResourceRecord));
          rd->name=strdup(q->qName);
          rd->type=A_Resource_RecordType;
          rd->rd_length=4;
          rd->class=1;
          rdns=get_DNS_Record(rd->rd_data.a_record.addr,&rd->name);
          if(rdns<0)
          {
          free(rd->name);
          free(rd);
          }
          else{
          msg->authorities=rd;
          msg->nsCount+=1;
          //free(rd->name);
          //free(rd);
          }
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
        printf(" \ncheck result !\n");
        print_resource_record(rr);
        if (count==0)
        {
          //free(rr->name);
          //free(rr);
          printf("cannot find %s",q->qName);
          struct ResourceRecord* rd=(struct ResourceRecord*)malloc(sizeof(struct ResourceRecord));
          rd->name=strdup(q->qName);
          rd->type=A_Resource_RecordType;
          rd->rd_length=4;
          rd->class=1;
          rdns=get_DNS_Record(rd->rd_data.a_record.addr,&rd->name);
          if(rdns<0)
          {
          free(rd->name);
          free(rd);
          }
          else{
          msg->authorities=rd;
          msg->nsCount+=1;
          }
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
        printf(" \ncheck result !\n");
        print_resource_record(rr);
        if (count==0)
        {
          struct ResourceRecord* rd=(struct ResourceRecord*)malloc(sizeof(struct ResourceRecord));
          rd->name=strdup(q->qName);
          rd->type=A_Resource_RecordType;
          rd->rd_length=4;
          rd->class=1;
          rdns=get_DNS_Record(rd->rd_data.a_record.addr,&rd->name);
          if(rdns<0)
          {
          free(rd->name);
          free(rd);
          }
          else{
          msg->authorities=rd;
          msg->nsCount+=1;
          //free(rd->name);
          //free(rd);
          }
          free(rr->name);
          free(rr);
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
  print_resource_record(msg->answers);
  //msg->msg_length=msg->msg_length+ResourceField_Caculate(msg);
  
  return resolve_result;
}

/* @return 0 upon failure, 1 upon success */
int encode_resource_records(struct ResourceRecord* rr, uint8_t** buffer)
{
  int i;
  printf("\n\n encode below \n\n");
  print_resource_record(rr);
  while (rr)
  {
    printf("\nencoding rr\n");
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
    printf("\n put in success\n");
    rr = rr->next;
  }

  return 0;
}

/* @return 0 upon failure, 1 upon success */
int encode_msg(struct Message* msg, uint8_t** buffer)
{
  struct Question* q;
  int rc;

  encode_header(msg, buffer);

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

void free_resource_records(struct ResourceRecord* rr)
{
  struct ResourceRecord* next;

  while (rr) {
    free(rr->name);
    next = rr->next;
    free(rr);
    rr = next;
  }
}

void free_questions(struct Question* qq)
{
  struct Question* next;

  while (qq) {
    free(qq->qName);
    next = qq->next;
    free(qq);
    qq = next;
  }
}

int main()
{
  // buffer for input/output binary packet
  uint8_t buffer[BUF_SIZE];
  struct sockaddr_in client_addr;
  socklen_t addr_len = sizeof(struct sockaddr_in);
  struct sockaddr_in addr;
  int nbytes, rc;
  int sock;
  int port = 53;

  struct Message msg;
  memset(&msg, 0, sizeof(struct Message));

  addr.sin_family = AF_INET;
  addr.sin_addr.s_addr = inet_addr(LOCAL_UADDR);
  addr.sin_port = htons(port);

  sock = socket(AF_INET, SOCK_DGRAM, 0);
  int on=1;
  if((setsockopt(sock,SOL_SOCKET,SO_REUSEADDR,&on,sizeof(on)))<0){
    printf("Set sock error:%s\n",strerror(errno));
  }
  printf("Set success\n");
  rc = bind(sock, (struct sockaddr*) &addr, addr_len);

  if (rc != 0)
  {
    printf("Could not bind: %s\n", strerror(errno));
    return 1;
  }

  printf("Listening on port %u.\n", port);

  while (1)
  {
    free_questions(msg.questions);
    free_resource_records(msg.answers);
    free_resource_records(msg.authorities);
    free_resource_records(msg.additionals);
    memset(&msg, 0, sizeof(struct Message));

    nbytes = recvfrom(sock, buffer, sizeof(buffer), 0, (struct sockaddr *) &client_addr, &addr_len);

    if (decode_msg(&msg, buffer, nbytes) != 0) {
      continue;
    }

    /* Print query */
    print_query(&msg);

    resolver_process(&msg);

    /* Print response */
    print_query(&msg);

    uint8_t *p = buffer;
    if (encode_msg(&msg, &p) != 0) {
      continue;
    }

    int buflen = p - buffer;
    sendto(sock, buffer, buflen, 0, (struct sockaddr*) &client_addr, addr_len);
  }
}
