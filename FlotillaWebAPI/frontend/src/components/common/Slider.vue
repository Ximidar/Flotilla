<template>
    <div class="organizer">
        <div class="slider-item">{{Name}}:</div>
        <div class="slidecontainer">
            <input type="range" 
                   :min="MinSize" 
                   :max="MaxSize" 
                   v-model="value" 
                   @input="changeValue" 
                   class="slider" 
                   >        
        </div>
        <div class="slider-item">{{value}}</div>
    </div>
</template>

<script>
export default {
    name: "Slider",
    props:["MaxSize", "MinSize", "Name", "InitialValue"],
    data() {
        return{
            value: 0
        }
    },
    created(){
        if (!isNaN(this.InitialValue)){
            this.value = this.InitialValue
            return
        }
        this.value = this.MinSize
        
    },
    methods:{
        changeValue(){
            this.$emit('SliderUpdate', this.value)
        }
    },
}
</script>

<style>

.organizer{
    display: flex;
    flex-direction: row;
    align-content: center;
    justify-content: center;
    align-items: center;
}

.slider-item{
    padding: 10px;
}

.slider {
    -webkit-appearance: none;  /* Override default CSS styles */
    appearance: none;
    height: 25px; /* Specified height */
    background: #146fa3; /* Grey background */
    outline: none; /* Remove outline */
    opacity: 0.7; /* Set transparency (for mouse-over effects on hover) */
    -webkit-transition: .2s; /* 0.2 seconds transition on hover */
    transition: opacity .2s;
}

/* Mouse-over effects */
.slider:hover {
    opacity: 1; /* Fully shown on mouse-over */
}

/* The slider handle (use -webkit- (Chrome, Opera, Safari, Edge) and -moz- (Firefox) to override default look) */ 
.slider::-webkit-slider-thumb {
    -webkit-appearance: none; /* Override default look */
    appearance: none;
    width: 25px; /* Set a specific slider handle width */
    height: 25px; /* Slider handle height */
    background: rgb(0, 253, 55); /* Green background */
    cursor: pointer; /* Cursor on hover */
}

.slider::-moz-range-thumb {
    width: 25px; /* Set a specific slider handle width */
    height: 25px; /* Slider handle height */
    background: #4CAF50; /* Green background */
    cursor: pointer; /* Cursor on hover */
}
</style>

