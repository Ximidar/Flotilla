<template>
  <div class="files">
    <div class="path-bar">
      PATH
    </div>
    <div class="file-items">
      <ul v-for="file in Contents" v-bind:key="file.Path">
        <FileItem :FileName="file.Name"
                  :Path="file.Path"
                  :PreviousPath="file.PreviousPath"
                  :FileType="file.FileType"
                  :Size="file.Size"
                  :UnixTime="file.UnixTime">
        </FileItem>
      </ul>  
    </div>
  </div>
</template>

<script>
var proto_file = require('./js_proto/FileStructures_pb')
import FileItem from '@/views/Files/FileItem'
export default {
  name: 'FlotillaFiles',
  components:{
    FileItem
  },
  data(){
    return{
      FileList: {},
      Contents: [],
      TestFileList: {
        "Path": "/tmp/testing/FileManager",
        "FileType": "folder",
        "Contents": [
            {
                "PreviousPath": "/tmp/testing/FileManager",
                "Name": "3D_Benchy.gcode",
                "Path": "/tmp/testing/FileManager/3D_Benchy.gcode",
                "FileType": "file",
                "Size": 4878333,
                "UnixTime": 1566708663
            },
            {
                "PreviousPath": "/tmp/testing/FileManager",
                "Name": "3D_Relative_Benchy.gcode",
                "Path": "/tmp/testing/FileManager/3D_Relative_Benchy.gcode",
                "FileType": "file",
                "Size": 4878333,
                "UnixTime": 1566708663
            },
            {
                "PreviousPath": "/tmp/testing/FileManager",
                "Name": "test",
                "Path": "/tmp/testing/FileManager/test",
                "FileType": "folder",
                "Size": 4096,
                "IsDir": true,
                "UnixTime": 1566708663
            }
        ]
      },
    }
  },
  
  methods:{
    Get_Files: function(){
      var root = new proto_file.File()
      for (var key in this.TestFileList){
        if (key == "Path"){
          root.Path = this.TestFileList[key]
        }
        if (key == "FileType"){
          root.FileType = this.TestFileList[key]
        }
        if (key == "Contents"){
          for (var item in this.TestFileList["Contents"]){
            var tempFile = new proto_file.File()
            tempFile.Path = this.TestFileList["Contents"][item]["Path"]
            tempFile.PreviousPath = this.TestFileList["Contents"][item]["PreviousPath"]
            tempFile.Name = this.TestFileList["Contents"][item]["Name"]
            tempFile.UnixTime = this.TestFileList["Contents"][item]["UnixTime"]
            tempFile.Size = this.TestFileList["Contents"][item]["Size"]
            tempFile.FileType = this.TestFileList["Contents"][item]["FileType"]
            if (tempFile.FileType == "folder"){
              tempFile.IsDir = this.TestFileList["Contents"][item]["IsDir"]
            } else {
              tempFile.IsDir = false
            }
            if (!root.Contents){
              root.Contents = []
            }
            root.Contents.push(tempFile)
          }
        }
      }
      this.FileList = root
    },
  ProcessFileList: function(){
    for (var file in this.FileList.Contents){
      this.Contents.push(this.FileList.Contents[file])
    }
  }
},
created(){
  this.Get_Files()
  this.ProcessFileList()
}

}
</script>

<style>
.files{
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: White;
}

.path-bar {
  text-align: left;
  border-bottom: 2px solid white ;
  overflow: hidden;
  padding-top: 10px;
  padding-bottom: 10px;
  
}

.file-items{
  text-align: left;
  overflow: hidden;
  padding-top: 10px;

}

ul{
  width: 100%;
  margin:0;
  padding:0;
}
</style>