<template>
  <v-container>
    <v-toolbar dense color="primary">
      <v-btn icon v-on:click.native="GoPrevious">
        <v-icon>$vuetify.icons.solid_arrow_left</v-icon>
      </v-btn>
      <v-toolbar-title class="px-3" >{{ FileList.CurrentFL.Path }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn icon>
        <v-icon>$vuetify.icons.solid_plus</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>$vuetify.icons.solid_folder_plus</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>$vuetify.icons.solid_ellipsis_h</v-icon>
      </v-btn>
    </v-toolbar>
    <v-container>
          <FileItem v-for="file in Contents" v-bind:key="file.Path"
            :File=file
            @clicked="ClickFile">
          </FileItem>
    </v-container>
  </v-container>
</template>

<script>
import FileItem from '@/views/Files/FileItem'
import FileInfo from '@/views/Files/FileInfo'
import ArrowLeft from '@/assets/svg/solid/arrow-left.svg'
import Vue from 'vue'
import { isNullOrUndefined } from 'util'
import flotilla from "@/flotilla"

export default {
  name: 'FlotillaFiles',
  components:{
    FileItem,
    FileInfo,
    ArrowLeft
  },
  mixins: [flotilla],
  data(){
    return{
      RootFS: {},
      FileList: {
        CurrentFL: {},
        PreviousFL: null,
      },
      Contents: [],
      SelectedFile: {}
    }
  },
  
  methods:{
  ClickFile: function(file){
    if (file.IsDir){
      console.log("Switching to ", file.Name)
      this.SwitchTo(file)
    } else {
      console.log("Selecting File")
      this.SelectedFile = file
    }
  },
  RequestFiles: function(){
    
    this.flotGetFiles().then( (files) => {
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
  width:40%;
  height: 80%;
  float: left;
  /* border: 2px solid blue ; */

}

.file-info{
  text-align: left;
  overflow: hidden;
  padding-top:10px;
  width:40%;
  height: 80%;
  /* border: 2px solid red ; */
}

.filecontainer{
  width: 80%;
  padding: 10px;
  margin:auto;
  overflow: auto;


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