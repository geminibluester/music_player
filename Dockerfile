FROM alpine:latest
WORKDIR /root/
COPY music_player .
RUN chmod +x /root/music_player
EXPOSE 8080
CMD [ "/root/music_player"]