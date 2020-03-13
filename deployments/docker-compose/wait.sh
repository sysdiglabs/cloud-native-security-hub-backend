apk update && apk add inotify-tools
echo "waiting for file..."
file=/shared-resources/file.unlock
while [ ! -f "$file" ]
do
  inotifywait -qqt 10 -e create -e moved_to "$(dirname $file)"
done
echo "found file"