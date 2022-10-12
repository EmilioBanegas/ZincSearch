import os
import re
import os 
import json
import base64
import requests

#Config ZincSearch
index = "correos"
zinc_server = "http://localhost:4080/api"
zinc_uid = "admin"
zinc_pwd = "Complexpass#123"
#Path DB
folder = os.path.join(os.path.dirname(os.path.abspath("./")), "db/enron_mail_20110402/maildir") 


bas64encoded_creds = base64.b64encode(bytes(zinc_uid + ":" + zinc_pwd, "utf-8")).decode("utf-8")
headers = {"Content-type": "application/json", "Authorization": "Basic " + bas64encoded_creds}

# Making Requets.
def sendRequest(typeR ,url,  headers, data, printR = True):
  if (typeR=="get"):
    resp = requests.get(url, headers = headers, data = data)
  elif (typeR == "post"):
    resp = requests.post(url, headers = headers, data = data)
  elif (typeR == "put"):
    resp = requests.put(url, headers = headers, data = data)
  else:
    resp = requests.delete(url, headers = headers, data = data)
  
  if printR:
    print(typeR, url)
    print(resp)
    print()  
  elif resp.status_code!=200:
    print(typeR, url)
    print(resp)
    print()
    
#Deleting index
def deleteIndex():
  sendRequest("delete", url = "{0}/index/{1}".format(zinc_server, index), headers = headers,  data=json.dumps({}))

#Creating index
def createIndex():
  body = {
    "name": index,
    "storage_type": "disk",
    "shard_num": 3,
    "mappings": {
      "properties": {
        "_index": {
          "type": "keyword",
          "index": True,
          "sortable": True,
          "aggregatable": True
        },
        "Message-ID": {
          "type": "keyword",
          "index": True,
          "store": True,
          "highlightable": True  
        },
        "From": {
          "type": "keyword",
          "index": True,
          "store": True,
          "highlightable": True  
        },
        "To": {
          "type": "text",
          "index": True,
          "store": True,
          "highlightable": True  
        },
        "Subject": {
          "type": "text",
          "index": True,
          "store": True,
          "highlightable": True  
        },
        "Bcc": {
          "type": "text",
          "index": True,
          "store": True,
          "highlightable": True  
        },
        "Cc": {
          "type": "text",
          "index": True,
          "store": True,
          "highlightable": True  
        },
        "Content": {
          "type": "text",
          "index": True,
          "store": True,
          "highlightable": True
        }
      }
    },

    "analysis": {
      "analyzer": {
        "default": {
          "type": "standard"
        },
        "my_analyzer": {
          "tokenizer": "standard",
          "char_filter": [
            "my_mappings_char_filter"
          ]
        }
      },
      "char_filter": {
        "my_mappings_char_filter": {
          "type": "mapping",
          "mappings": [
            "&lt; => _<_",
            "&gt; => _>_"
          ]
        }
      }
    }
  }

  sendRequest("post", "{0}/index".format(zinc_server), headers = headers, data =  json.dumps(body))

def insertData():
    
  #Obteniendo el directorio de cada usuario
  f = []
  for (root,dirs,files) in os.walk(folder, topdown=True):
    f.extend(dirs)
    break

  #Recorriendo los subdirectorios de cada usuario para crear el objeto JSON
  for user in f:
    for (root, dirs, files) in os.walk(os.path.join(folder, user), topdown = True):
      strPosition = re.search(user, root)
      stringDir = root[strPosition.end()+1:]
      beforeDirs = stringDir.split("\\")
      for fileS in files:
        #reading file that contain email
        f = open(root+"/"+fileS, mode = 'r', encoding= 'windows-1252')
        lines = f.readlines()
        f.close()

        #creating object dictionary that represent email for current file readed
        emailData = dict()

        #converting file to email
        lastKey = ''

        n = len(lines)
        for i in range(n):
          x = re.search("^[a-zA-Z-]+:", lines[i])
          if(x == None and lastKey == "X-FileName"):
            emailData[lastKey] = re.sub("(^(\n|\s)+|(\n|\s)+$)","", emailData[lastKey], 2)
            emailData["Content"] = re.sub("(^(\n|\s)+|(\n|\s)+$)","", "".join(lines[i:n]), 2) 
            break

          if(x == None and lastKey!= ""):
            emailData[lastKey] += lines[i] 
          elif(x != None):
            if(lastKey != ""):
              emailData[lastKey] = re.sub("(^(\n|\s)+|(\n|\s)+$)","", emailData[lastKey], 2)
            lastKey = lines[i][x.start(): (x.end()-1)]
            emailData[lastKey] = lines[i][len(lastKey)+1:]
        
        #index email to DB into value of index
        sendRequest(typeR= "post", url = "{0}/{1}/_doc".format(zinc_server, index), headers = headers, data = json.dumps(emailData), printR = False)

      print(user, stringDir)

def main():
  deleteIndex()
  createIndex() 
  insertData()

if __name__ == "__main__":
  main()