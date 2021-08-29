<template>
<div class="settings">
    <Slider Name="Size Width" MaxSize="4" MinSize="1" @SliderUpdate="changeWidth"/>
    <Slider Name="Size Height" MaxSize="4" MinSize="1" @SliderUpdate="changeHeight"/>
</div>
    
</template>

<script>
import Slider from "@/components/common/Slider"

export default {
    name: 'WidgetSettings',
    data() {
        return {
            MaxSize: 4,
            SizeW: 1,
            SizeH: 1,
        }
    },
    props: {
        WidgetName: {
            default: 'none',
            type: String
        },
    },
    methods: {
        changeWidth(newWidth){
            this.SizeW = newWidth
            var updateCoord = {
                Name: this.WidgetName,
                Coord: this.SizeW,
                elementID: 'w'
            }
            this.$store.commit("updateWidgetCoord", updateCoord)
        },
        changeHeight(newHeight){
           this.SizeH = newHeight
            var updateCoord = {
                Name: this.WidgetName,
                Coord: this.SizeH,
                elementID: 'h'
            }
            this.$store.commit("updateWidgetCoord", updateCoord)
        }
    },
    components:{
        Slider
    }    
}
</script>

<style scoped>

.settings{
    border: 2px solid red;
    color: white;
    display: flex;
    flex-direction: column;
    align-content: center;
    justify-content: center;
    align-items: center;
}
</style>

