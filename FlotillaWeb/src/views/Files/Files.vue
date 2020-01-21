<template>
  <div class="files">
    <div class="path-bar">
      <ArrowLeft v-on:click="GoPrevious" class="iconsize" />
      {{ FileList.CurrentFL.Path }}
    </div>
    <div class="file-items">
      <ul v-for="file in Contents" v-bind:key="file.Path">
        <FileItem :File=file
                  @clicked="ClickFile">
        </FileItem>
      </ul>  
    </div>
  </div>
</template>

<script>
import FileItem from '@/views/Files/FileItem'
import ArrowLeft from '@/assets/svg/solid/arrow-left.svg'
import Vue from 'vue'
import { isNullOrUndefined } from 'util'
import { Flotilla } from "@/flotilla"

export default {
  name: 'FlotillaFiles',
  components:{
    FileItem,
    ArrowLeft
  },
  data(){
    return{
      RootFS: {},
      FileList: {
        CurrentFL: {},
        PreviousFL: null,
      },
      Contents: [],
    }
  },
  
  methods:{
  ClickFile: function(file){
    if (file.IsDir){
      console.log("Switching to ", file.Name)
      this.SwitchTo(file)
    } else {
      console.log(file.Name)
    }
  },
  RequestFiles: function(){
    console.log(Flotilla)
    var flot = new Flotilla()
    flot.GetFiles().then( (files) => {
      this.RootFS = files
      this.SwitchTo(this.RootFS)
    })
  },
  GoPrevious: function(){
    if (!isNullOrUndefined(this.FileList.PreviousFL) && this.FileList.PreviousFL != {}){
      console.log("Defined Previous", this.FileList.PreviousFL)
      this.FileList = this.FileList.PreviousFL
      this.ProcessCurrentFL()
    } else {
      this.SwitchTo(this.RootFS)
    }
  },
  SwitchTo: function(file){
    this.Contents = []
    console.log("Switching to", file.Path)
    if (file.Path != this.RootFS.Path){
      this.FileList.PreviousFL = JSON.parse(JSON.stringify(this.FileList))
    } else {
      this.FileList.PreviousFL = null
    }
    this.FileList.CurrentFL = file
    this.ProcessCurrentFL()
  },
  ProcessCurrentFL: function(){
    this.Contents = []
    for (var file in this.FileList.CurrentFL.Contents){
      this.Contents.push(this.FileList.CurrentFL.Contents[file])
    }
  }
},
created(){
  this.RequestFiles()
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

.iconsize{
  width: 30px;
  height: 30px;
  fill:white;
} 

</style>