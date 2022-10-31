<template>
  <div class="container">
  	<div class="page-header">
       <h1 class="page-title">Laporan Clock</h1>
    </div>
    <div class="row row-cards">
      <div class="col-12">
      <div class="card">
        <div class="card-header">
          <div class="col-2" style="padding-left: 0px;">
            <date-picker v-model="bulan" value-type="YYYY-MM" format="MM/YYYY" type="month" placeholder="Pilih Bulan" @input="onBulan()"></date-picker>
          </div>
          <div class="card-options">
            <button @click="downloadExcel()" class="btn btn-primary" :disabled="disableexport"><i class="fe fe-file-text"></i> Export Excel</button>
          </div>
        </div>
        <div class="card-body" style="padding:0 !important;">
          <div v-if="bulan === null" class="jadwal-null">
            <h2 class="text-muted"> Silahkan Pilih Bulan </h2>
          </div>
          <div v-else class="laporan-clock">           
            <div class="laporan-clock-title">
              <h3>TABEL KEHADIRAN GURU</h3>
              <div class="tahunbulan">
                <h3 class="bulan">{{moment(bulan).format('MMMM')}}</h3>
                <h6 class="tahun">TAHUN {{moment(this.bulan).format('YYYY')}}</h6>
              </div>
            </div>
            <div v-if="clockData.length === 0" class="jadwal-null" style="background-color:#fff;border-top-style:solid;border-top-color:#e8e8e8;border-top-width: 0.01em;">
                <h2 class="text-muted"> Data Kosong</h2>         
            </div>
            <div v-else class="laporan-clock-body">
              <div class="laporan-clock-user" v-for="user in clockData" :key="user.user_id">
                <div class="clock-nama">
                  {{user.mapel_guru}}
                </div>
                <div class="clock-tanggal">
                  <div v-for="indexTgl in 31" :key="indexTgl">
                    {{indexTgl}}
                  </div>
                </div>
                <div class="clock-item">
                  <div v-for="index in 31" :key="index" class="item">
                    <div v-if="user.tgl.includes(`${index}`)">
                    <div v-for="(val,key) in user.mulai[user.tgl.indexOf(`${index}`)].split(',')" :key="key" style="margin-bottom:8px">                      
                      <div style="background: lightgray;">
                        {{val}}
                      </div>
                      <div style="background: orange;">
                        {{user.keluar[user.tgl.indexOf(`${index}`)].split(',')[key]}}
                      </div>
                    </div>  
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="card-footer" style="padding: 20px;"></div>
      </div>
      </div>      
    </div>
  </div>
</template>
<script type="application/javascript">

import moment from "moment";

export default { 
  components: {
    
  },
  data () {
    return {
      moment:moment,
      disableexport:true,
      bulan:null,
      clockData:[],
    }
  },
  watch: {
    clockData(){
      if(this.clockData.length == 0){
        this.disableexport = true;
      }else{
        this.disableexport = false;
      }
    }
  },
  methods: {
    /** Method **/
    onBulan(){
      if(this.bulan != null){
        this.fetchData();
      }
    },
    downloadExcel(){
      this.disableexport = true;
      axios.post(window.location.origin + "/api/laporan/clock",{
        bulan:moment(this.bulan).format('MMMM'),tanggal:this.bulan,data:JSON.stringify(this.clockData)
      }).then(response => {
            var download = response.data.data;
            window.open(window.location.origin+"/"+download);
            this.disableexport = false;
      }).catch(error => { this.disableexport = false; })
    },
    fetchData() {
      axios.get(window.location.origin + "/api/laporan/clock?tanggal="+this.bulan)
           .then(response => {               
            this.AbsensiList(response.data.data);
      }).catch(error => this.clockData = [])
    },  
    AbsensiList(newArr) { 
      this.clockData = newArr.reduce(function(o, cur) {

        // Get the index of the key-value pair.
        var occurs = o.reduce(function(n, item, i) {
          return (item.user_id === cur.user_id) ? i : n;
        }, -1);

        // If the name is found,
        if (occurs >= 0) {

          // append the current value to its list of values.          
          o[occurs].tgl = o[occurs].tgl.concat(cur.tgl);
          o[occurs].mulai = o[occurs].mulai.concat(cur.mulai);
          o[occurs].keluar = o[occurs].keluar.concat(cur.keluar);

          // Otherwise,
        } else {

          // add the current item to o (but make sure the value is an array).
          var obj = {
            user_id:cur.user_id,            
            mapel_guru: cur.mapel_guru,            
            tgl: [cur.tgl],
            mulai: [cur.mulai],
            keluar: [cur.keluar]
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