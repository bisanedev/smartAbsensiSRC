<template>
<modal name="absen" :adaptive="true" :width="820" height="100%" @before-open="beforeOpen" @before-close="beforeClose">
<div class="card">
  <div class="card-header" style="padding: 0rem !important;">
      <h3 style="position: absolute;margin: auto;text-align: center;left: 0;right: 0;">
        Edit Absensi {{tanggal | formatDate}}
      </h3>
      <div class="card-options">
       <button class="close" @click="keluarModal" style="margin-right: 20px;">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
  </div>
  <div class="card-body card-absen">      
    <div class="card-overflow col-md-12">
        <div class="siswacards" v-for="(data, key1) in sortedArray" :key="data.id">
        <div class="row">
          <div v-if="data.foto" class="avatar d-block avfoto" v-bind:style="{ 'background-image':'url(./storage/siswa/'+ data.foto +')' }">
            <div class="urut-absen">{{ data.urut }}</div>
          </div>
          <div v-else-if="data.jenis == 'Perempuan'" class="avatar d-block avfoto" style="background-image: url(./assets/images/cewek.png)">
            <div class="urut-absen">{{ data.urut }}</div>
          </div>
          <div v-else class="avatar d-block avfoto" style="background-image: url(./assets/images/cowok.png)">
            <div class="urut-absen">{{ data.urut }}</div>
          </div>          
          <div class="col-md-8" style="margin-top: 17px;">
            <b class="namaAbsen">{{ data.nama.toLowerCase() }}</b>
            <p style="font-size: 12px;line-height: 15px !important;margin-bottom:3px !important;">
            {{ data.jenis }}
            </p>
            <p style="font-size:10px;font-weight:bold;">{{data.mapel_guru}} [{{data.mapel}}] ~ {{moment(data.mapel_start).format('HH:mm')}}-{{moment(data.mapel_end).format('HH:mm')}}
            </p>
          </div>
          <div v-if="data.kode == 'h'" class="pilih" v-bind:style="{ 'background-color':'green' }"> 
          <div class="item-action dropdown">
            <a href="javascript:void(0)" data-toggle="dropdown" style="font-weight: bold;color:white;cursor: pointer;">Hadir</a>
                <div class="dropdown-menu dropdown-menu-right">
                <button type="button" @click="changeKode('h',data.id)" class="dropdown-item" > 
                <p style="color:green">  Hadir </p>
                </button>
                <button type="button" @click="changeKode('i',data.id)" class="dropdown-item" >
                <p style="color:blue">  Izin </p> 
                </button>
                <button type="button" @click="changeKode('s',data.id)" class="dropdown-item" >
                <p style="color:#07ade3">  Sakit </p> 
                </button>
                <button type="button" @click="changeKode('t',data.id)" class="dropdown-item" >
                <p style="color:orange">  Terlambat </p>  
                </button>
                <button type="button" @click="changeKode('a',data.id)" class="dropdown-item" >
                 <p style="color:red">  Alpa </p>  
                </button>
                </div>
          </div>             
          </div>
          <div v-if="data.kode == 'i'" class="pilih" v-bind:style="{ 'background-color':'blue' }"> 
          <div class="item-action dropdown">
            <a href="javascript:void(0)" data-toggle="dropdown" style="font-weight: bold;color:white;cursor: pointer;">Izin</a>
                <div class="dropdown-menu dropdown-menu-right">
                <button type="button" @click="changeKode('h',data.id)" class="dropdown-item" > 
                <p style="color:green">  Hadir </p>
                </button>
                <button type="button" @click="changeKode('i',data.id)" class="dropdown-item" >
                <p style="color:blue">  Izin </p> 
                </button>
                <button type="button" @click="changeKode('s',data.id)" class="dropdown-item" >
                <p style="color:#07ade3">  Sakit </p> 
                </button>
                <button type="button" @click="changeKode('t',data.id)" class="dropdown-item" >
                <p style="color:orange">  Terlambat </p>  
                </button>
                <button type="button" @click="changeKode('a',data.id)" class="dropdown-item" >
                 <p style="color:red">  Alpa </p>  
                </button>
                </div>
          </div> 
          </div>
          <div v-if="data.kode == 's'" class="pilih" v-bind:style="{ 'background-color':'#07ade3' }"> 
          <div class="item-action dropdown">
            <a href="javascript:void(0)" data-toggle="dropdown" style="font-weight: bold;color:white;cursor: pointer;">Sakit</a>
                <div class="dropdown-menu dropdown-menu-right">
                <button type="button" @click="changeKode('h',data.id)" class="dropdown-item" > 
                <p style="color:green">  Hadir </p>
                </button>
                <button type="button" @click="changeKode('i',data.id)" class="dropdown-item" >
                <p style="color:blue">  Izin </p> 
                </button>
                <button type="button" @click="changeKode('s',data.id)" class="dropdown-item" >
                <p style="color:#07ade3">  Sakit </p> 
                </button>
                <button type="button" @click="changeKode('t',data.id)" class="dropdown-item" >
                <p style="color:orange">  Terlambat </p>  
                </button>
                <button type="button" @click="changeKode('a',data.id)" class="dropdown-item" >
                 <p style="color:red">  Alpa </p>  
                </button>
                </div>
          </div>           
          </div>
          <div v-if="data.kode == 't'" class="pilih" v-bind:style="{ 'background-color':'orange' }">          
          <div class="item-action dropdown">
            <a href="javascript:void(0)" data-toggle="dropdown" style="font-weight: bold;color:white;cursor: pointer;">Terlambat</a>
                <div class="dropdown-menu dropdown-menu-right">
                <button type="button" @click="changeKode('h',data.id)" class="dropdown-item" > 
                <p style="color:green">  Hadir </p>
                </button>
                <button type="button" @click="changeKode('i',data.id)" class="dropdown-item" >
                <p style="color:blue">  Izin </p> 
                </button>
                <button type="button" @click="changeKode('s',data.id)" class="dropdown-item" >
                <p style="color:#07ade3">  Sakit </p> 
                </button>
                <button type="button" @click="changeKode('t',data.id)" class="dropdown-item" >
                <p style="color:orange">  Terlambat </p>  
                </button>
                <button type="button" @click="changeKode('a',data.id)" class="dropdown-item" >
                 <p style="color:red">  Alpa </p>  
                </button>
                </div>
          </div>               
          </div>
          <div v-if="data.kode == 'a'" class="pilih" v-bind:style="{ 'background-color':'red' }"> 
          <div class="item-action dropdown">
            <a href="javascript:void(0)" data-toggle="dropdown" style="font-weight: bold;color:white;cursor: pointer;">Alpa</a>
                <div class="dropdown-menu dropdown-menu-right">
                <button type="button" @click="changeKode('h',data.id)" class="dropdown-item" > 
                <p style="color:green">  Hadir </p>
                </button>
                <button type="button" @click="changeKode('i',data.id)" class="dropdown-item" >
                <p style="color:blue">  Izin </p> 
                </button>
                <button type="button" @click="changeKode('s',data.id)" class="dropdown-item" >
                <p style="color:#07ade3">  Sakit </p> 
                </button>
                <button type="button" @click="changeKode('t',data.id)" class="dropdown-item" >
                <p style="color:orange">  Terlambat </p>  
                </button>
                <button type="button" @click="changeKode('a',data.id)" class="dropdown-item" >
                 <p style="color:red">  Alpa </p>  
                </button>
                </div>
          </div>            
          </div>
        </div>
        <div v-if="data.kode == 's' || data.kode == 'i'" style="margin-left: -13px;background-color: #f2f1f1;">
            <input type="text" class="form-control" :value="data.kode_note" 
            style="border-radius: 0px !important;width: 81.5% !important;" @change="changeNote($event.target.value,data.id)" placeholder="Keterangan">
        </div>
      </div>
    </div>
  </div>
  <div class="bottomfixed">
    <div class="col-md-12" style="align-items:center;justify-content:center;display:flex;">
    <button style="width: 250px;" class="btn btn-primary" @click="updateSiswa()">
      Update Absensi
    </button>
    </div>
  </div>
</div>
</modal>
</template>
<script type="application/javascript">
import moment from 'moment';
export default {
	name: 'absen',
  	data () {
  		return {
        moment: moment,
        disableclose:true,
        absen:[],
        siswa_id:'',
        tanggal:'',
      }
  },
  computed: {
    sortedArray: function() {

      function compare(a, b) {
        if (a.jam < b.jam)
          return -1;
        if (a.jam > b.jam)
          return 1;
        return 0;
      }

      return this.absen.sort(compare);
    }
  },
  methods: {
    changeNote(ket,id){
       let i = this.absen.findIndex(r => r.id === id);    
       this.$set(this.absen[i], 'kode_note', `${ket}`);
    },
    changeKode(kode,id){
       let i = this.absen.findIndex(r => r.id === id);    
       this.$set(this.absen[i], 'kode', `${kode}`);
       this.$set(this.absen[i], 'kode_note', '');
    },
    fetchAbsen(){

    axios.get(window.location.origin+"/api/absensi/getabsen?tanggal="+moment(this.tanggal).format('YYYY-MM-DD')+"&siswa="+this.siswa_id)
    .then(response => { 
          this.absen = response.data.data          
    }).catch(error => this.absen = [])

    },
    updateSiswa:function(event){
      axios.post(window.location.origin+"/api/absensi/getabsen/store", {
        absen:JSON.stringify(this.absen)          
        })
        .then(response => {
          this.$parent.fetchAbsen();
          if(response.data.status == "success")
          { 
            //console.log('berhasil diupdate');
            //this.disableclose = false;
            //this.$modal.hide('siswa')
            this.$notify({
                group: 'notifikasi',                  
                title: 'Update Absensi Sukses',
                type: 'success',
                text: 'Data Absensi Berhasil Di Update'
            });
            
          }
        })
        .catch(error => {
        
            //console.log('gagal');
            this.$notify({
              group: 'notifikasi',                  
              title: 'Update Gagal',
              type: 'error',
              text: error.response.data.message
            });      
          
        })  
    },
    keluarModal(){
      this.disableclose = false;      
      this.$modal.hide('absen');      
    },
	  beforeOpen (event) { 
      this.disableclose = true;
      this.siswa_id = event.params.siswaid;
      this.tanggal = event.params.tgl;
      this.fetchAbsen();
    },
    beforeClose (event) {
      //console.log("close")
      if (this.disableclose){
        event.stop();
      }   
    }
  }
  //----------
}   
</script>