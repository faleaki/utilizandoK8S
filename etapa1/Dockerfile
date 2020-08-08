FROM nginx:1.19.0-alpine
RUN rm /etc/nginx/conf.d/default.conf
RUN mkdir -p /var/www/html
COPY index.html /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
