FROM scratch

WORKDIR /carefree

COPY door .

COPY project/door/frontend/build /carefree/project/door/frontend/build

ENTRYPOINT [ "./door" ]
