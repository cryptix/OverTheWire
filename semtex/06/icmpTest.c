#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <netinet/in.h>
#include <netinet/in_systm.h>
#include <netinet/ip.h>
#include <netinet/ip_icmp.h>
#include <string.h>
#include <arpa/inet.h>
 
#define BUFFER_SIZE 400
#define PACKET_DELAY_USEC 30
 
char buf[BUFFER_SIZE];
 

void set_ip_layer_fields(struct icmphdr *icmp, struct ip *ip)
{
    // IP Layer
    ip->ip_v = 4;
    ip->ip_tos = 0;
    ip->ip_len = htons(sizeof(buf));
    ip->ip_id = htons(4321);
    ip->ip_off = htons(0);
    ip->ip_ttl = 255;
    ip->ip_p = 1;
    ip->ip_hl = sizeof*ip >> 2;
    ip->ip_sum = 0; /* Let kernel fill in */
 
    // ICMP Layer
    icmp->type = ICMP_ECHO;
    icmp->code = 0;	
    icmp->checksum = htons(~(ICMP_ECHO << 8));	
}
 
void set_socket_options(int s)
{
    int on = 1;
 
    // socket options, tell the kernel we provide the IP structure 
    if(setsockopt(s, IPPROTO_IP, IP_HDRINCL, &on, sizeof(on)) < 0){
        perror("setsockopt() for IP_HDRINCL error");
        exit(1);
    }	
}
 
int main(int argc, char *argv[])
{
    int s;	
    struct ip *ip = (struct ip *)buf;
    struct icmphdr *icmp = (struct icmphdr *)(ip + 1);
    struct sockaddr_in dst;

    if (argc == 2) {
        // Clear data paylod
        memset(buf, 0, sizeof(buf));

        // s = argv[1];
        // Create RAW socket 
        if((s = socket(AF_INET, SOCK_RAW, IPPROTO_RAW)) < 0){
            perror("socket() error");
            exit(1);
        }

        set_socket_options(s);

        ip->ip_dst.s_addr = inet_addr("134.28.179.196");
        ip->ip_src.s_addr = inet_addr("188.138.121.9");

        set_ip_layer_fields(icmp, ip);

        dst.sin_addr = ip->ip_dst;
        dst.sin_family = AF_INET;

        if(sendto(s, buf, sizeof(buf), 0, (struct sockaddr *)&dst, sizeof(dst)) < 0){
            fprintf(stderr, "Error during packet send.\n");
            perror("sendto() error");
        }else
            printf("sendto() is OK.\n");

        close(s);
    }
    return 0;
}