# Updating an expiring certificate
## Prerequisites
Before starting, make a backup of the existing (expiring) certificate and private key:
1. ssh into the server
2. Enter the directory where the certificates are stored: `cd /root/deployment/mnt/config/nginx-conf/conf.d`
3. Backup the current private key by copying it: `cp server.key server.key.old`
    - The current private key is now the file `server.key.old`.
4. Backup the current certificate: `cp singlishwords.nus.edu.sg.chained.crt singlishwords.nus.edu.sg.chained.crt.old`
    - The current certificate is now the file `singlishwords.nus.edu.sg.chained.crt.old`.

Then, delete the existing certificate and private key:
1. Ensure that you are in the directory where the certificates are stored by following steps 1 and 2 above
2. `rm server.key` to delete the private key
3. `rm singlishwords.nus.edu.sg.chained.crt` to delete the certificate

## Downloading a new certificate
1. Go to (https://ncertrequest.nus.edu.sg/). Ensure that you are connected to NUS VPN. 
2. On the top navigation bar, click on "Dashboards", then "All Certificates"
3. Click on "My Certificates"
4. Under nicknames, click on `singlishwords.nus.edu.sg`
    - If it cannot be found, you will need to go back to the All Certificates Dashboard and look in 'Certificates Expiring within 30 Days', then renew the certificate on NCertRequest before continuing
5. Click on "Actions", the spanner icon in the top-right-hand corner. Click "Download".
6. On the pop-up dialog box, ensure that the format is PEM (OpenSSL). 
7. Checkboxes for "Root Chain" and "Private Key" should already have been checked; if they are not, check them. Leave Chain Order as "End Entity First".
8. The password fields are for you to set a password. Set a password different from other passwords used for the site. 
9. Below, check "Extract PEM content into separate files (.crt, .key)".
10. Click "Download". `_.nus.edu.sg.zip`, should have been downloaded to your computer. 

## Installing the new certificate
1. Unzip `_.nus.edu.sg.zip`. You will need a few files:
    * `_.nus.edu.sg-chain.pem` (chained, intermediate certificates by the certificate authority)
    * `_.nus.edu.sg.crt` (the actual certificate for the server)
    * `_.nus.edu.sg.key` (password protected private key)
2. Remove the password from the private key. See [here](https://www.madboa.com/geek/openssl/#how-do-i-remove-a-passphrase-from-a-key) for more details
    * On the UNIX command line, `openssl rsa -in _.nus.edu.sg.key -out server.key`
    * You will be prompted to enter the password for the private key. See step 8 under the previous section. 
    * The private key to be copied to the server later will be the file `server.key`. 
3. Chain the intermediate certs and the server cert. See [here](https://nginx.org/en/docs/http/configuring_https_servers.html#chains).
    * The server cert must appear before the intermediate certs
    * On the UNIX command line, `cat _.nus.edu.sg.crt _.nus.edu.sg-chain.pem > singlishwords.nus.edu.sg.chained.crt`
    * The chained certificate to be copied to the server later will be the file `singlishwords.nus.edu.sg.chained.crt`.
4. Copy the certificate and key to the appropriate directory in the server. 
    ```
     scp -r ./server.key root@172.105.127.13:/root/deployment/mnt/config/nginx-conf/conf.d
     scp -r ./singlishwords.nus.edu.sg.chained.crt root@172.105.127.13:/root/deployment/mnt/config/nginx-conf/conf.d
    ```
    * Alternatively, if you are not comfortable with the command line, a program like WinSCP can also be used.
5. ssh into the server, then enter the certificate directory.
    ```bash
     ssh root@172.105.127.13
     cd /root/deployment/mnt/config/nginx-conf/conf.d
    ```
6. Make sure the cert and key is readable by everyone (if not, nginx cannot read them):
    * `ls -l`
    * Check the permissions on the left, there should be 3 'r's. For example: 
    ```
    [root@localhost conf.d]# ls -l
    total 40
    -rw-r--r-- 1 swosw swosw 1715 Nov  4 16:48 default.conf
    -rw-r--r-- 1 root  root  1679 Nov  4 16:56 server.key
    -rw-r--r-- 1 root  root  1679 Nov  4 16:37 server.key.new
    -rw-r--r-- 1 root  root  1674 Nov  4 06:18 server.key.old
    -rw-r--r-- 1 root  root  5860 Nov  4 16:56 singlishwords.nus.edu.sg.chained.crt
    -rw-r--r-- 1 root  root  5860 Nov  4 16:32 singlishwords.nus.edu.sg.chained.crt.new
    -rw-r--r-- 1 root  root  4305 Nov  4 06:18 singlishwords.nus.edu.sg.chained.crt.old
    ```
    * If new permissions need to be added, do `chmod a+r <FILENAME>` where <FILENAME> is the name of the file. This adds permissions for everyone to read the file ([link](https://kb.iu.edu/d/abdb)).
7. Navigate back to the `deployment` parent folder: `cd /root/deployment` 
8. Restart the nginx container
   ```
   docker-compose build nginx
   docker-compose stop nginx
   docker-compose up -d --no-deps nginx
   ```
9. Check that the container does not have any errors
  ```
  docker ps
  ```
10. Finally, check that the certificate is actually updated on the website itself. 
    * Go to `https://singlishwords.nus.edu.sg/`
    * Example: On Google Chrome, click on the lock icon in the address bar, then "Connection is secure", then "Certificate is valid". On the certificate viewer that pops up, the validity period can be seen. 