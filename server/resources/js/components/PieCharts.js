import {Pie} from 'vue-chartjs'

export default {
  extends: Pie,
  props:['data'],
  mounted() {
    this.renderPieChart();
  },
  computed: {
    chartData: function() {
      return this.data;
    }
  },
  methods:{
    renderPieChart:function(){
        this.renderChart({
        labels: ['Hadir', 'Izin', 'Sakit', 'Terlambat','Alpa'],      
        datasets: [
        {
          backgroundColor: [
            'rgba(10, 152, 11, .8)',
            'rgba(10, 31, 253, .8)',
            'rgba(10, 168, 253, .8)',
            'rgba(255, 191,0, .8)',
            'rgba(255, 0,0, .8)'
          ],
          borderWidth: 0,
          data: [this.chartData.hadir,this.chartData.izin,this.chartData.sakit,this.chartData.terlambat,this.chartData.alpa]
        }
        ]
        }, {
          responsive: true,
          maintainAspectRatio: false,
          tooltips: {
            enabled: true,
            callbacks: {
              label:function (tooltipItem, data) {
                let dataset = data.datasets[tooltipItem.datasetIndex]
                let currentValue = dataset.data[tooltipItem.index]
                let total = dataset.data.reduce(function (previousValue, currentValue, currentIndex, array) {
                  return previousValue + currentValue
                })
                return Math.floor(((currentValue / total) * 100) + 0.5) + '%'
              }
            }
          }
        })      
    }
  //end of method
  },
  watch: {
    data: function() {
      this.$data._chart.destroy();      
      this.renderPieChart();
    }
  }
//end of default
}