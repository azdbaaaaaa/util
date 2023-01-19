export name="bill_server"
cp -a configmap.template.yaml out.yaml
file=`cat out.yaml`
printf "`export -p`\ncat << EOF\n$file\nEOF" | bash > out.yaml
