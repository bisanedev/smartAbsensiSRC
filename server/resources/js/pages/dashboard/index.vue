<template>
  <div class="container">
  	<div class="page-header">
       <h1 class="page-title">Dashboard</h1>
    </div>
    <div class="row row-cards">
      <div class="col-6 col-sm-4 col-lg-2">
        <div class="card">
          <div class="card-body p-3 text-center">
            <div class="h1 m-0" style="line-height: 1.6;font-size: 3rem;color:#868e96;">{{jumlahsiswa-(countJumlah('h')+countJumlah('t')+countJumlah('i')+countJumlah('s')+countJumlah('a'))}}</div>
            <div class="text-muted mb-4">Siswa Tidak Absensi</div>
          </div>
        </div>
      </div>
      <div class="col-6 col-sm-4 col-lg-2">
        <div class="card">
          <div class="card-body p-3 text-center">
            <div class="h1 m-0" style="line-height: 1.6;font-size: 3rem;color:#7bd235;">{{countJumlah('h')}}</div>
            <div class="text-muted mb-4" style="color:#7bd235 !important;">Siswa Hadir</div>
          </div>
        </div>
      </div>
      <div class="col-6 col-sm-4 col-lg-2">
        <div class="card">
          <div class="card-body p-3 text-center">
          <div class="h1 m-0" style="line-height: 1.6;font-size: 3rem;color:#07ade3;">{{countJumlah('s')}}</div>
            <div class="text-muted mb-4"><a style="color:#07ade3;" href="javascript:void(0)" @click="showModal(ListDetail('s'),'Siswa Sakit')">Siswa Sakit</a> <i class="fe fe-more-vertical" style="font-size: 12px;"></i></div>
          </div>
        </div>
      </div>
      <div class="col-6 col-sm-4 col-lg-2">
        <div class="card">
          <div class="card-body p-3 text-center">
          <div class="h1 m-0" style="line-height: 1.6;font-size: 3rem;color:blue;">{{countJumlah('i')}}</div>
            <div class="text-muted mb-4"><a style="color:blue;" href="javascript:void(0)" @click="showModal(ListDetail('i'),'Siswa Izin')">Siswa Izin</a> <i class="fe fe-more-vertical" style="font-size: 12px;"></i></div>
          </div>
        </div>
      </div>
      <div class="col-6 col-sm-4 col-lg-2">
        <div class="card">
          <div class="card-body p-3 text-center">
          <div class="h1 m-0" style="line-height: 1.6;font-size: 3rem;color:orange;">{{countJumlah('t')}}</div>
            <div class="text-muted mb-4"><a style="color:orange;" href="javascript:void(0)" @click="showModal(ListDetail('t'),'Siswa Terlambat')">Siswa Terlambat</a> <i class="fe fe-more-vertical" style="font-size: 12px;"></i></div>
          </div>
        </div>
      </div>
      <div class="col-6 col-sm-4 col-lg-2">
        <div class="card">
          <div class="card-body p-3 text-center">
            <div class="h1 m-0" style="line-height: 1.6;font-size: 3rem;color:#cd201f;">{{countJumlah('a')}}
            </div>
            <div class="text-muted mb-4"><a style="color:#cd201f;" href="javascript:void(0)" @click="showModal(ListDetail('a'),'Siswa Alpa')">Siswa Alpa</a> <i class="fe fe-more-vertical" style="font-size: 12px;"></i></div>
          </div>
        </div>
      </div>
      <div class="col-6 col-lg-6">
        <div class="card">
        <div class="card-header">
          <h3 class="card-title">Hari {{moment().format('dddd')}}</h3>
        </div>
        <div class="card-body">
          <pie-charts :data="sekarang"></pie-charts>
        </div>
        </div>
      </div>     
      <div class="col-6 col-lg-6">
        <div class="card">
        <div class="card-header">          
          <div style="padding-left: 0px;width:218px;">
            <v-select :options="SemesterList" label="nama" :reduce="nama => nama.waktustart" v-model="semester" placeholder="Pilih Semester" @input="onPick($event)"></v-select>
          </div>
        </div>
        <div class="card-body">
          <pie-charts :data="semesterData"></pie-charts>
        </div>
        </div>
      </div>       
    </div>
    <detail-siswa-modal/>
  </div>
</template>
<script type="application/javascript">
import PieCharts from "../../components/PieCharts";
import DetailSiswaModal from './DetailSiswaModal.vue'
import moment from "moment";
import "moment/locale/id";

export default { 
  components: {
    PieCharts,
    DetailSiswaModal
  },
  data () {
    return {
      moment: moment,
      jumlahsiswa:'',
      absen:[],
      semester:null,      
      SemesterList:[],      
      semesterData:"",
      sekarang:""              
    }
  },
  created() {
    this.fetchSemester();
    this.RenderWebsocket();
    this.fetchChartsNow();
    this.fecthAbsen();
  },
  methods: {
    /** Detail Siswa **/
    showModal(data,title){
      this.$modal.show('detail-siswa',{wow:data,header:title});
    },
    //list detail
    ListDetail(kode){
      var data = this.absen; 
      return data.filter(item => item.kode === kode )      
    },
    //count jumlah
    countJumlah(kode){
      var data = this.absen; 
      return data.filter(item => item.kode === kode ).length      
    },
    /** ONPick ***/
    onPick(event) {
      if (this.semester != null) {
        this.fetchSemesterTgl();
      } 
    },
    fecthAbsen(){
        axios.get(window.location.origin + "/api/info/absen/harian")
        .then(response => {
          this.jumlahsiswa = response.data.jumlah
          this.absen = response.data.absen
        }).catch(error => this.absen = this.jumlahsiswa = [])
    },
    fetchSemester() {
        axios.get(window.location.origin + "/api/semester/lists")
        .then(response => {
          this.SemesterList = response.data.data          
        }).catch(error => this.SemesterList = [])
    },
    /** websocket **/
    RenderWebsocket: function() {
        this.socket = new WebSocket("ws://"+window.location.host+"/websocket/absen/ws");
        this.socket.onopen = () => {
          console.log("trigger Konek");                     
            this.socket.onmessage = ({data}) => {          
            this.fetchChartsNow();
            this.fecthAbsen();
          };
        }
        this.socket2 = new WebSocket("ws://"+window.location.host+"/websocket/absenupdate/ws");
        this.socket2.onopen = () => {
          console.log("Realtime Update Konek");                     
            this.socket2.onmessage = ({data}) => {          
            this.fetchChartsNow();
            this.fecthAbsen();
          };
        }
    },
    //get jumlah moment(this.tanggal).format('YYYY-MM-DD')
    fetchChartsNow() {
        axios.get(window.location.origin + "/api/info/siswa/harian")
        .then(response => {
          this.sekarang = response.data.jumlah          
        }).catch(error => this.sekarang = [])
    },
    fetchSemesterTgl() {
      if(this.semester == null){return}
        axios.get(window.location.origin + "/api/info/siswa/semester?tanggal="+moment(this.semester).format('YYYY-MM-DDTHH:mm'))
        .then(response => {
          this.semesterData = response.data.jumlah          
        }).catch(error => this.semesterData = [])
    },
    //eend
  }  
  //----
}
</script>