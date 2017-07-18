Using dev mode
==============

Normally chaincodes are started and maintained by peer. However in “dev
mode", chaincode is built and started by the user. This mode is useful
during chaincode development phase for rapid code/build/run/debug cycle
turnaround.

We start "dev mode" by leveraging pre-generated orderer and channel artifacts for
a sample dev network.  As such, the user can immediately jump into the process
of compiling chaincode and driving calls.


Download docker images
^^^^^^^^^^^^^^^^^^^^^^

We need four docker images in order for "dev mode" to run against the supplied
docker compose script.  If you installed the ``fabric-samples`` repo clone and
followed the instructions to :ref:`download-platform-specific-binaries`, then
you should have the necessary Docker images installed locally.

.. note:: If you choose to manually pull the images then you must retag them as
          ``latest``.

Issue a ``docker images`` command to reveal your local Docker Registry.  You
should see something similar to following:

.. code:: bash

  docker images
  REPOSITORY                     TAG                                  IMAGE ID            CREATED             SIZE
  hyperledger/fabric-tools       latest                               e09f38f8928d        4 hours ago         1.32 GB
  hyperledger/fabric-tools       x86_64-1.0.0-rc1-snapshot-f20846c6   e09f38f8928d        4 hours ago         1.32 GB
  hyperledger/fabric-orderer     latest                               0df93ba35a25        4 hours ago         179 MB
  hyperledger/fabric-orderer     x86_64-1.0.0-rc1-snapshot-f20846c6   0df93ba35a25        4 hours ago         179 MB
  hyperledger/fabric-peer        latest                               533aec3f5a01        4 hours ago         182 MB
  hyperledger/fabric-peer        x86_64-1.0.0-rc1-snapshot-f20846c6   533aec3f5a01        4 hours ago         182 MB
  hyperledger/fabric-ccenv       latest                               4b70698a71d3        4 hours ago         1.29 GB
  hyperledger/fabric-ccenv       x86_64-1.0.0-rc1-snapshot-f20846c6   4b70698a71d3        4 hours ago         1.29 GB

.. note:: If you retrieved the images through the :ref:`download-platform-specific-binaries`,
          then you will see additional images listed.  However, we are only concerned with
          these four.

Now open three terminals and navigate to your ``chaincode-docker-devmode``
directory in each.

Terminal 1 - Start the network
------------------------------

.. code:: bash

    docker-compose -f docker-compose-simple.yaml up

The above starts the network with the ``SingleSampleMSPSolo`` orderer profile and
launches the peer in "dev mode".  It also launches two additional containers -
one for the chaincode environment and a CLI to interact with the chaincode.  The
commands for create and join channel are embedded in the CLI container, so we
can jump immediately to the chaincode calls.

Terminal 2 - Build & start the chaincode
----------------------------------------

.. code:: bash

  docker exec -it chaincode bash

You should see the following:

.. code:: bash

  root@d2629980e76b:/opt/gopath/src/chaincode#

Now, compile your chaincode:

.. code:: bash

  cd chaincode_example02
  go build

Now run the chaincode:

.. code:: bash

  CORE_PEER_ADDRESS=peer:7051 CORE_CHAINCODE_ID_NAME=mycc:0 ./chaincode_example02

The chaincode is started with peer and chaincode logs indicating successful registration with the peer.
Note that at this stage the chaincode is not associated with any channel. This is done in subsequent steps
using the ``instantiate`` command.

Terminal 3 - Use the chaincode
------------------------------

Even though you are in ``--peer-chaincodedev`` mode, you still have to install the
chaincode so the life-cycle system chaincode can go through its checks normally.
This requirement may be removed in future when in ``--peer-chaincodedev`` mode.

We'll leverage the CLI container to drive these calls.

.. code:: bash

  docker exec -it cli bash

.. code:: bash

  peer chaincode install -p chaincodedev/chaincode/chaincode_example02 -n mycc -v 0
  peer chaincode instantiate -n mycc -v 0 -c '{"Args":["init","a","100","b","200"]}' -C myc

Now issue an invoke to move ``10`` from ``a`` to ``b``.

.. code:: bash

  peer chaincode invoke -n mycc -c '{"Args":["invoke","a","b","10"]}' -C myc

Finally, query ``a``.  We should see a value of ``90``.

.. code:: bash

  peer chaincode query -n mycc -c '{"Args":["query","a"]}' -C myc

Testing new chaincode
---------------------

By default, we mount only ``chaincode_example02``.  However, you can easily test different
chaincodes by adding them to the ``chaincode`` subdirectory and relaunching
your network.  At this point they will be accessible in your ``chaincode`` container.

.. Licensed under Creative Commons Attribution 4.0 International License
     https://creativecommons.org/licenses/by/4.0/


peer chaincode install -p chaincodedev/chaincode/sacc -n mycc -v 0
peer chaincode instantiate -n mycc -v 0 -c '{"Args":["a","10"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["set", "a", "20"]}' -C myc
peer chaincode query -n mycc -c '{"Args":["query","a"]}' -C myc

----------------------------
| Create Loan Application  |  
----------------------------

INSTALL :
  
  peer chaincode install -p chaincodedev/chaincode/marbles02 -n mycc -v 0

INSTANTIATE :
  
  peer chaincode instantiate -n loanApp -v 2 -c '{"Args":["initMarble","marble1","blue","35","tom"]}' -C myc  

INVOKE  :

  peer chaincode invoke -n loanApp -c '{"Args":["CreateLoanApplication"]}' -C myc


Marbles :

INSTALL :


  // ====CHAINCODE EXECUTION SAMPLES (CLI) ==================





peer chaincode install -p chaincodedev/chaincode/marbles02 -n mycctv -v 0


peer chaincode instantiate -n mycctv -v 0 -c '{"Args":[]}' -C myc


peer chaincode invoke -C myc -n mycctv -v 0 -c '{"Args":["initMarble","marble1","blue","35","tom"]}'
peer chaincode invoke -C myc -n mycctv -v 0 -c '{"Args":["initMarble","marble2","red","50","tom"]}'
peer chaincode invoke -C myc -n mycctv -v 0 -c '{"Args":["initMarble","marble3","blue","70","tom"]}'

peer chaincode invoke -C myc -n mycctv -v 0 -c '{"Args":["transferMarble","marble2","jerry"]}'

peer chaincode invoke -C myc -n mycctv -v 0 -c '{"Args":["transferMarblesBasedOnColor","blue","jerry"]}'

peer chaincode invoke -C myc -n mycctv -v 0 -c '{"Args":["delete","marble1"]}'


peer chaincode query -C myc -n mycctv -v 0 -c '{"Args":["readMarble","marble1"]}'
peer chaincode query -C myc -n mycctv -v 0 -c '{"Args":["readMarble","marble3"]}'

peer chaincode query -C myc -n mycctv -v 0 -c '{"Args":["getMarblesByRange","marble1","marble3"]}'
peer chaincode query -C myc -n mycctv -v 0 -c '{"Args":["getHistoryForMarble","marble1"]}'

// Rich Query (Only supported if CouchDB is used as state database):
peer chaincode query -C myc -n mycctv -v 0 -c '{"Args":["queryMarblesByOwner","tom"]}'
peer chaincode query -C myc -n mycctv -v 0 -c '{"Args":["queryMarbles","{\"selector\":{\"owner\":\"tom\"}}"]}'


// ====CHAINCODE EXECUTION SAMPLES (CLI) ==================

// ==== Invoke marbles ====
// peer chaincode invoke -C myc1 -n marbles -c '{"Args":["initMarble","marble1","blue","35","tom"]}'
// peer chaincode invoke -C myc1 -n marbles -c '{"Args":["initMarble","marble2","red","50","tom"]}'
// peer chaincode invoke -C myc1 -n marbles -c '{"Args":["initMarble","marble3","blue","70","tom"]}'

// peer chaincode invoke -C myc1 -n marbles -c '{"Args":["transferMarble","marble2","jerry"]}'
// peer chaincode invoke -C myc1 -n marbles -c '{"Args":["transferMarblesBasedOnColor","blue","jerry"]}'
// peer chaincode invoke -C myc1 -n marbles -c '{"Args":["delete","marble1"]}'

// ==== Query marbles ====
// peer chaincode query -C myc1 -n marbles -c '{"Args":["readMarble","marble1"]}'
// peer chaincode query -C myc1 -n marbles -c '{"Args":["getMarblesByRange","marble1","marble3"]}'
// peer chaincode query -C myc1 -n marbles -c '{"Args":["getHistoryForMarble","marble1"]}'

// Rich Query (Only supported if CouchDB is used as state database):
//   peer chaincode query -C myc1 -n marbles -c '{"Args":["queryMarblesByOwner","tom"]}'
//   peer chaincode query -C myc1 -n marbles -c '{"Args":["queryMarbles","{\"selector\":{\"owner\":\"tom\"}}"]}'

//The following examples demonstrate creating indexes on CouchDB
//Example hostname:port configurations
//
//Docker or vagrant environments:
// http://couchdb:5984/
//
//Inside couchdb docker container
// http://127.0.0.1:5984/

// Index for chaincodeid, docType, owner.
// Note that docType and owner fields must be prefixed with the "data" wrapper
// chaincodeid must be added for all queries
//
// Definition for use with Fauxton interface
// {"index":{"fields":["chaincodeid","data.docType","data.owner"]},"ddoc":"indexOwnerDoc", "name":"indexOwner","type":"json"}
//
// example curl definition for use with command line
// curl -i -X POST -H "Content-Type: application/json" -d "{\"index\":{\"fields\":[\"chaincodeid\",\"data.docType\",\"data.owner\"]},\"name\":\"indexOwner\",\"ddoc\":\"indexOwnerDoc\",\"type\":\"json\"}" http://hostname:port/myc1/_index
//

// Index for chaincodeid, docType, owner, size (descending order).
// Note that docType, owner and size fields must be prefixed with the "data" wrapper
// chaincodeid must be added for all queries
//
// Definition for use with Fauxton interface
// {"index":{"fields":[{"data.size":"desc"},{"chaincodeid":"desc"},{"data.docType":"desc"},{"data.owner":"desc"}]},"ddoc":"indexSizeSortDoc", "name":"indexSizeSortDesc","type":"json"}
//
// example curl definition for use with command line
// curl -i -X POST -H "Content-Type: application/json" -d "{\"index\":{\"fields\":[{\"data.size\":\"desc\"},{\"chaincodeid\":\"desc\"},{\"data.docType\":\"desc\"},{\"data.owner\":\"desc\"}]},\"ddoc\":\"indexSizeSortDoc\", \"name\":\"indexSizeSortDesc\",\"type\":\"json\"}" http://hostname:port/myc1/_index

// Rich Query with index design doc and index name specified (Only supported if CouchDB is used as state database):
//   peer chaincode query -C myc1 -n marbles -c '{"Args":["queryMarbles","{\"selector\":{\"docType\":\"marble\",\"owner\":\"tom\"}, \"use_index\":[\"_design/indexOwnerDoc\", \"indexOwner\"]}"]}'

// Rich Query with index design doc specified only (Only supported if CouchDB is used as state database):
//   peer chaincode query -C myc1 -n marbles -c '{"Args":["queryMarbles","{\"selector\":{\"docType\":{\"$eq\":\"marble\"},\"owner\":{\"$eq\":\"tom\"},\"size\":{\"$gt\":0}},\"fields\":[\"docType\",\"owner\",\"size\"],\"sort\":[{\"size\":\"desc\"}],\"use_index\":\"_design/indexSizeSortDoc\"}"]}'


HISTORY :

[  
   {  
      "TxId":"44ba571f1ff85bd48c19b5b6199cd30c7b30dae46b4a0b4d42c4cdb1ee04ae88",
      "Value":{  
         "docType":"marble",
         "name":"marble1",
         "color":"blue",
         "size":35,
         "owner":"tom"
      },
      "Timestamp":"2017-07-18 08:01:50.023626055 +0000 UTC",
      "IsDelete":"false"
   },
   {  
      "TxId":"1564d2b7b00f08ee294845886e1f290211b6391efd9403389b471884b0442583",
      "Value":{  
         "docType":"marble",
         "name":"marble1",
         "color":"blue",
         "size":35,
         "owner":"jerry"
      },
      "Timestamp":"2017-07-18 08:24:12.364205332 +0000 UTC",
      "IsDelete":"false"
   },
   {  
      "TxId":"7e887e52bdfe0bf0842cc5908686e6ee332ad94328ecad4a9f82d1342fb71d2a",
      "Value":{  
         "docType":"marble",
         "name":"marble1",
         "color":"blue",
         "size":35,
         "owner":"mahendra"
      },
      "Timestamp":"2017-07-18 08:26:23.668568412 +0000 UTC",
      "IsDelete":"false"
   }
]
