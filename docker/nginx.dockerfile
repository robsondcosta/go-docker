FROM nginx:latest 
COPY /frontend /var/www/public 
COPY /docker/config/nginx.conf /etc/nginx/nginx.conf 
RUN chmod 755 -R /var/www/public 
ENTRYPOINT ["nginx"]

CMD ["-g","daemon off;"]

# necessário para uma aplicação frontend