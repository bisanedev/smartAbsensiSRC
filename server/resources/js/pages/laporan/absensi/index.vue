<template>
  <div class="container">
  	<div class="page-header">
       <h1 class="page-title">Laporan Absensi</h1>
    </div>
    <div class="row row-cards row-deck">          
       <div class="col-12">
        <div class="card">
          <div class="card-header">
            <div class="col-2" style="padding-left: 0px;">
             <v-select :options="kelasData" label="nama" :reduce="nama => nama.id" v-model="kelas" placeholder="Pilih Kelas" @input="onKelas()"></v-select>
            </div>
            <div class="col-2" style="padding-left: 0px;">
             <date-picker v-model="bulan" value-type="YYYY-MM" format="MM/YYYY" type="month" placeholder="Pilih Bulan" @input="onBulan()" :disabled="disableBulan"></date-picker>
            </div>
            <div class="card-options">
              <button @click="downloadExcel()" class="btn btn-primary" :disabled="disableexport"><i class="fe fe-file-text"></i> Export Excel</button>
            </div>
          </div>
          <div class="card-body" style="padding:0 !important;">
            <div v-if="bulan === null" class="jadwal-null">
              <h2 class="text-muted"> Silahkan Pilih Kelas & Pilih Bulan </h2>
            </div>
            <div v-else class="laporan-absensi">
              <div class="laporan-header">
                <div class="nama">Nama Siswa</div>
                <div class="hari">
                  <div v-for="index in 31" :key="index" class="item">{{index}}</div>
                </div>
              </div>
              <div v-if="siswaData.length === 0" class="laporan-null text-muted">
                <div></div>
                <div>Data Kosong</div>                
              </div>
              <div v-else v-for="data in siswaData" :key="data.urut" class="data-laporan">
                <div class="data-nama"><p class="no">{{data.urut | leftPad}}.</p><p class="nama">{{data.nama.substring(0,34).toLowerCase()}}</p></div>
                <div class="data-hari">
                  <div v-for="index in 31" :key="index" class="data-item">                    
                    <a v-if="data.tgl.includes(`${index}`)" href="javascript:void(0)" v-bind:style="{ 'color':colorAbsen(data.tgl.indexOf(`${index}`),data.kode) }" @click="changeAbsen(data.tgl[data.tgl.indexOf(`${index}`)],data.siswa_id)">
                      {{cekAbsenSempurna(data.tgl.indexOf(`${index}`),data.kode)}}
                    </a>
                    <div v-else></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="card-footer" style="padding: 20px;"></div>
        </div> 
       </div>        
    </div>
    <absen-modal/>
  </div>
</template>
<script type="application/javascript">

import moment from "moment";
import "moment/locale/id";
import AbsenModal from "./AbsenModal.vue";

export default { 
  components: {
    AbsenModal    
  },
  data () {
    return {
      disableexport:true,
      disableBulan:true,
      moment:moment,
      kelas:null,
      kelasName:'',
      bulan:null,
      kelasData:[],
      siswaData:[]      
    }
  },
  created() {
   this.fetchKelas();   
  },  
  watch: {
    kelas(newVal) {
      if(newVal != null){        
        this.disableBulan=false;        
      }else{
        this.disableBulan=true;
        this.bulan=null;
      }
    },
    siswaData(){
      if(this.siswaData.length == 0){
        this.disableexport = true;
      }else{
        this.disableexport = false;
      }
    }
  },
  computed: {
       
  },
  methods: {
    /** Method **/
    onBulan(){      
      this.fetchAbsen();      
    },
    onKelas(){
      if(this.bulan != null){
        this.fetchAbsen();
      }
      this.getKelasAttr();
    },
    getKelasAttr(){
      for (var key in this.kelasData) {
          if (this.kelasData.hasOwnProperty(key)) {
            if(this.kelas == this.kelasData[key].id)
              {                
                this.kelasName = this.kelasData[key].nama;                  
              }
          }
      }
    },
    downloadExcel(){
      axios.post(window.location.origin + "/api/laporan/kelas/absensi",{
        data:JSON.stringify(this.siswaData),bulan:moment(this.bulan).format('MMMM'),kelas:this.kelasName,tgl:this.bulan
      }).then(response => {
            var download = response.data.data;
            window.open(window.location.origin+"/"+download);            
      }).catch(error => "Gagal")
    },
    changeAbsen(tanggal,siswa_id) {
      var date = this.bulan+"-"+tanggal;    
      this.$modal.show('absen',{tgl:moment(date).format('YYYY-MM-DD'),siswaid:siswa_id});      
    },
    colorAbsen(index,kode){
      var absenArr = kode[index].split(",");
      var jumlah = absenArr.length;      
      var hadir = absenArr.filter(function(absen){ return absen === "h" || absen === "t"}).length;      
      var sakit = absenArr.filter(function(absen){ return absen === "s"}).length;
      var izin = absenArr.filter(function(absen){ return absen === "i"}).length;
      var alpa = absenArr.filter(function(absen){ return absen === "a"}).length;
      if (jumlah === hadir){return "green";}
      else if (jumlah === sakit){return "#07ade3";}
      else if (jumlah === izin){return "blue";}
      else if (jumlah === alpa){return "red";}
      else {return "orange";}      
    },
    cekAbsenSempurna(index,kode){      
      var absenArr = kode[index].split(",");
      var jumlah = absenArr.length;      
      var hadir = absenArr.filter(function(absen){ return absen === "h" || absen === "t"}).length;      
      var sakit = absenArr.filter(function(absen){ return absen === "s"}).length;
      var izin = absenArr.filter(function(absen){ return absen === "i"}).length;
      var alpa = absenArr.filter(function(absen){ return absen === "a"}).length;
      if (jumlah === hadir){return "H";}
      else if (jumlah === sakit){return "S";}
      else if (jumlah === izin){return "I";}
      else if (jumlah === alpa){return "A";}
      else {return "!";}
    },
    fetchKelas() {
      axios.get(window.location.origin + "/api/allkelas")
           .then(response => {
            this.kelasData = response.data.data
      }).catch(error => this.kelasData = [])
    },
    fetchAbsen() {
      axios.get(window.location.origin + "/api/laporan/kelas/absensi?kelas="+this.kelas+"&tanggal="+this.bulan)
           .then(response => {            
            //this.siswaData = response.data.data;            
            this.AbsensiList(response.data.data);            
      }).catch(error => this.siswaData = [])
    },
    AbsensiList(newArr) {
      this.siswaData = newArr.reduce(function(o, cur) {

        // Get the index of the key-value pair.
        var occurs = o.reduce(function(n, item, i) {
          return (item.urut === cur.urut) ? i : n;
        }, -1);

        // If the name is found,
        if (occurs >= 0) {

          // append the current value to its list of values.
          o[occurs].tgl = o[occurs].tgl.concat(cur.tgl);
          o[occurs].kode = o[occurs].kode.concat(cur.kode);

          // Otherwise,
        } else {

          // add the current item to o (but make sure the value is an array).
          var obj = {
            siswa_id:cur.siswa_id,
            urut: cur.urut,
            nama: cur.nama,
            tgl: [cur.tgl],
            kode: [cur.kode]
          };
          o = o.concat([obj]);
        }

        return o;
      }, []);      
    }
    //eend
  }  
  //----
}
</script>