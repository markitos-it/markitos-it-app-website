FROM nginx:alpine

COPY index.html /usr/share/nginx/html/
COPY *.css /usr/share/nginx/html/
COPY *.js /usr/share/nginx/html/

EXPOSE 80

ENTRYPOINT ["nginx"]
CMD ["-g", "daemon off;"]