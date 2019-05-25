import Vue from 'vue'
import Vuex from 'vuex'

// Components
import PConsole from '@/components/widgets/pconsole'

Vue.use(Vuex)

function CoreWidget(name, svg, coord, widget){
  this.ItemName = name
  this.ItemSVGIcon = svg
  this.Coord = coord
  this.widget = widget
}

function Coord(x, y, w, h){
  this.x = x
  this.y = y
  this.w = w
  this.h = h
}

export default new Vuex.Store({
  state: {
    activeWidgets: [],
    coreWidgets: [
      new CoreWidget(
        'Motor Control',
        'svg1',
        new Coord(0,0,2,2),
        ''
      ),
      new CoreWidget(
        'PConsole',
        'svg1',
        new Coord(2,0,4,4),
        PConsole
      ),
      new CoreWidget(
        'Files',
        'svg1',
        new Coord(1,1,1,1),
        ''
      ),
      new CoreWidget(
        'File Progress',
        'svg1',
        new Coord(1,1,1,1),
        ''
      ),
      new CoreWidget(
        'Temperature Graph',
        'svg1',
        new Coord(1,1,1,1),
        ''
      ),
  ],
  SolarizedDark:{
    base03: '#002b36',
    base02: '#073642',
    base01: '#586e75',
    base00:  '#657b83',
    base0:   '#839496',
    base1:   '#93a1a1',
    base2:   '#eee8d5',
    base3:   '#fdf6e3',
    yellow:  '#b58900',
    orange:  '#cb4b16',
    red:     '#dc322f',
    magenta: '#d33682',
    violet:  '#6c71c4',
    blue:    '#268bd2',
    cyan:    '#2aa198',
    green:   '#859900',
  }
  },
  getters: {
    getCoreWidgets: state => {
      return state.coreWidgets
    },
  },
  mutations: {
    toggleActiveWidget (state, widgetName) {
      // If the name is already in the store then remove it, else add it

      var index = -1
      index = state.activeWidgets.findIndex(function(item){
        return item.name == widgetName
      })

      if (index !== -1){
        state.activeWidgets.splice(index, 1)
      } else {
        var index = state.coreWidgets.findIndex(function(item) {
          return item.ItemName == widgetName
        })

        state.activeWidgets.push({
          name: state.coreWidgets[index].ItemName,
          widget: state.coreWidgets[index].widget
        })
      }
      
    },
  updateWidgetCoord (state, updateCoord){
    var widgetName = updateCoord.Name
    var newCoord = updateCoord.Coord
    var elementID = updateCoord.elementID
    var index = state.coreWidgets.findIndex(function(item){
      return item.ItemName == widgetName
    })
    Vue.set(state.coreWidgets[index].Coord, elementID, newCoord)
  }
  },
  actions: {


  }
})
