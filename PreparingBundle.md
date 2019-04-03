Prepare Bundle for Manual Deployment

Static files to include:
* Executable Gopherface application
* StaticAssets


## Transpile the application
Create a tar ball

``` sh
$ cd $GOPHERCLUB_APP_ROOT/client
$ gopherjs build --verbose -m -o $GOPHERCLUB_APP_ROOT/static/js/client.min.js
```

We are now ready to create a deployment bundle, a tarball, which includes the igweb Linux binary along with the static folder. We create the tarball by issuing the following commands:
``` sh
cd $GOPHERCLUB_APP_ROOT
tar zcvf /tmp/bundle.tgz builds/gopherclub-linux64 static templates

scp /tmp/bundle.tgz gopherclub@gopherclub.jaynejacobs.com:/tmp/.
```

At this point, we now have a production deployment bundle that is ready to be deployed to the Linode instance.
163  mkdir ~/gopherclub
  164  cd ~/gopherclub/
  167  mv /tmp/bundle.tgz  .
  168  tar zxvf bundle.tgz 
  170  cd builds/
  171  ls
  mv builds/gopherclub-linux64 ./gopherclub
  172  ls -la
  173  cd ~/gopherclub/
  174  ls
  175  rm bundle.tgz 
