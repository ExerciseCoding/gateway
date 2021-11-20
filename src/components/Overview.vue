<template>
  <div id="myChart" :style="{width: '300px', height: '300px'}"></div>
</template>

<script>


export default {
  name: "Overview",
  data(){

  },
  mounted() {
    this.drawLine();
  },
  methods:{
    drawLine(){
      // 初始化echarts实例
      let myChart = this.$echarts.init(document.getElementById('myChart'))

      // 绘制图表
      myChart.setOption({
        title: {text: '请求QPS'},
        tooltip: {},
        xAxis: {
          data:['http','rpc','tcp']
        },
        yAxis:{},
        series:[
          {
            name:"QPS",
            type: 'bar',
            data: [76,23,52]
          }
        ]
      });
    }
  }
}
</script>

<style scoped>

</style>
