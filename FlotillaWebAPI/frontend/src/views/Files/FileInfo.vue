<template>
  <div class="info-block">
      <div class="file-name">{{ Name }}</div>
      <TextButton
        Label="Play"
        ButtonColor="green"
        :Action="PlayFile"
        class="play-button">
      </TextButton> 

  </div>
</template>

<script>
import TextButton from "@/components/common/TextButton"
import { Flotilla } from "@/flotilla"
export default {
    Name: "FileInfo",
    components: {
        TextButton
    },
    props:{
        File: {
            type: Object,
            default: {}
        }

    },
    data: function() {
        return {
            SelectedFile: {},
            Name: ""
        }
    },
    methods:{
        PlayFile: function(){
            console.log("playing file")
            let flot = new Flotilla()
            flot.SelectFile(this.SelectedFile).then(result => {
                console.log(result)
                flot.PostAction("Play")
            })
            
        }
        
    },
    watch: {
        File: function (newval, oldval){
            console.log("New File Selected")
            this.SelectedFile=newval
            this.Name=this.SelectedFile.Name
        }
    }

}
</script>

<style scoped>
.play-button{
    width: 80%;
}

.info-block{
  display:inline-block;
  width: 200px;
  height: 50px;
}

.file-name{
    height: 25px;
}


</style>