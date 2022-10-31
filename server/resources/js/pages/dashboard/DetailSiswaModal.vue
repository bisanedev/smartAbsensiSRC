<template>
<modal name="detail-siswa" height="auto" :scrollable="true" :width="600"  @before-open="beforeOpen" @before-close="beforeClose">
  <div class="card" style="margin-bottom: 0;">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;margin-left: 15px;">{{title}}</h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('detail-siswa')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">
      <div class="col-md-12" style="margin-top:20px;margin-bottom:20px;">        
        <div class="siswacards" v-for="(data, key1) in siswaList" :key="key1">
          <div class="row">
            <div v-if="data.foto" class="avatar d-block avfoto" v-bind:style="{ 'background-image':'url(./storage/siswa/'+ data.foto +')' }">            
            </div>
            <div v-else-if="data.jenis == 'Perempuan'" class="avatar d-block avfoto" style="background-image: url(./assets/images/cewek.png)">              
            </div>
            <div v-else class="avatar d-block avfoto" style="background-image: url(./assets/images/cowok.png)">
            </div>          
            <div class="col-md-8" style="margin-top: 10px;left:20px;">
              <b class="namaAbsen">{{ data.nama.toLowerCase() }}</b>
              <p style="font-size: 12px;line-height: 15px !important;margin-bottom:3px !important;">Kelas {{ data.kelas }}</p>
              <p style="font-size:10px;font-weight:bold;margin-bottom: 2px !important;">{{data.mapel_guru}} [{{data.mapel}}] ~ {{moment(data.mapel_start).format('HH:mm')}}-{{moment(data.mapel_end).format('HH:mm')}}
              </p>
              <p v-if="data.kode_note" style="font-size:10px;font-weight:bold;color:var(--mycolor);">"{{data.kode_note}}"</p>
            </div>
          </div>          
        </div>  
      </div>    
       
    </div>
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">
import moment from 'moment';
export default {
  name: 'DetailSiswaModal',  
  data () {
    return {
      moment: moment,
      siswaList:[], 
      title:'',
    }
  }, 
  methods: {        
    beforeOpen (event) {
      this.siswaList = event.params.wow;
      this.title = event.params.header;
    },
    beforeClose (event) {      
      this.siswaList=[];     
    }
   //end method
  }
}
</script>