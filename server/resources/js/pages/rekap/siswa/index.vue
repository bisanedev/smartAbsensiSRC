<template>
    <div class="container">
        <div class="page-header">
            <h1 class="page-title" style="white-space:pre;">Rekap Siswa </h1>
            <h3 class="page-title" v-if="bulan != null"> [{{moment(bulan).format('MMMM')}}]</h3>
            <div class="page-options d-flex">
                <date-picker v-model="bulan" :disabled-date="notAfterToday" value-type="YYYY-MM" format="MM/YYYY" type="month" placeholder="Pilih Bulan" @input="onBulan()"></date-picker>
            </div>
        </div>
        <div class="row row-cards">
            <div class="col-8">
                <div class="card">
                    <div class="card-header">
                        <h5 style="margin-bottom: 0 !important;">Persentase Kehadiran Siswa</h5>
                        <div class="card-options">
                        <select class="form-control custom-select w-auto" v-model="order" style="font-size: 0.7rem;font-weight:bold;" @change="onSort()">
                            <option value="asc">TERENDAH</option>
                            <option value="desc">TERTINGGI</option>
                        </select>
                        </div>
                    </div>
                    <div class="card-body rekapDataKosong" style="padding:0 !important;" v-if="tableData.length == 0">
                        <h3 class="text-muted"> Pilih Bulan / Data Kosong </h3>
                    </div>
                    <div v-else class="card-body" style="padding:0 !important;">
                        <div class="input-group">
                            <input type="text" class="form-control" v-model="search" v-on:keyup="searchEnter" placeholder="Cari Nama Siswa" style="border-radius: 0 !important;border: none;">
                            <button class="btn bg-transparent" @click="refresh()" style="margin-left: -40px; z-index: 100;"><i class="fe fe-x text-muted"></i></button>
                            <div class="input-group-append">
                                <button class="btn btn-primary" type="button" @click="cari()" style="border-radius: 0 !important;"><i class="fe fe-search"></i>
                                </button>
                            </div>
                        </div>
                        <div class="rekaplist">
                            <div class="rekaplist-item" v-for="(data,key) in tableData" :key="key">
                                <p style="margin-bottom:0!important;margin-top: 6px;text-transform:capitalize;width: 180px;">{{data.nama.substring(0,20).toLowerCase()}} [{{data.kelas}}] </p>
                                <div class="rekaplist-data">
                                    <div class="clearfix">
                                        <div class="float-left">
                                            <small class="text-muted">Tingkat Kehadiran</small>
                                        </div>
                                        <div class="float-right">
                                            <strong>{{percentageData(data)}}%</strong>
                                        </div>
                                    </div>
                                    <div class="progress progress-xs">
                                        <div class="progress-bar bg-green" role="progressbar" v-bind:style="{ 'width':percentageData(data)+'%' }" :aria-valuenow="percentageData(data)" aria-valuemin="0" aria-valuemax="100"></div>
                                    </div>
                                </div>
                                <button class="btn btn-primary" @click="rincian(data)" type="button" style="font-size:10px;font-weight: 600;line-height: 0;padding: 0.375rem 0.75rem"><i class="fe fe-user" style="font-size: 12px;"></i> RINCIAN</button>
                            </div>
                        </div>
                    </div>
                    <div class="card-footer">
                        <span v-if="pagination.total_record" style="position: absolute;bottom: 20px;"> &nbsp; <i>Total Data : {{ pagination.total_record }} entries.</i></span>
                        <nav v-if="pagination && tableData.length > 0">
                            <ul class="pagination float-right">
                                <li class="page-item" :class="{'disabled' : currentPage === 1}">
                                    <a class="page-link" href="#" @click.prevent="changePage(1)">
                                        <i class="fe fe-chevrons-left"></i>
                                    </a>
                                </li>
                                <li class="page-item" :class="{'disabled' : currentPage === 1}">
                                    <a class="page-link" href="#" @click.prevent="changePage(currentPage - 1)"><i class="fe fe-chevron-left"></i></a>
                                </li>
                                <li v-for="page in pagesNumber" v-if="Math.abs(page - currentPage) < 3 || page == pagesNumber - 1 || page == 0" class="page-item" :class="{'active': page == pagination.page}">
                                    <a href="javascript:void(0)" @click.prevent="changePage(page)" class="page-link">{{ page }}</a>
                                </li>
                                <li class="page-item" :class="{'disabled': currentPage === pagination.total_page}">
                                    <a class="page-link" href="#" @click.prevent="changePage(currentPage + 1)"><i class="fe fe-chevron-right"></i></a>
                                </li>
                                <li class="page-item" :class="{'disabled': currentPage === pagination.total_page}">
                                    <a class="page-link" href="#" @click.prevent="changePage(pagination.total_page)"><i class="fe fe-chevrons-right"></i></a>
                                </li>
                            </ul>
                        </nav>
                    </div>
                </div>
            </div>
            <div class="col-4">
                <div class="card">
                    <div class="card-header">
                        <h5 style="margin-bottom: 0 !important;">Rincian Kehadiran Siswa</h5>
                    </div>
                    <div class="card-body" style="padding:0 !important;">
                    <div v-if="rinciaData == ''" class="rincianDataKosong">
                            <h5 class="text-muted"> Pilih Rincian Siswa </h5>
                    </div>
                    <div v-else class="rincianData">
                        <div v-if="rinciaData.foto" class="avatar d-block rinciFoto" v-bind:style="{ 'background-image':'url(./storage/siswa/'+ rinciaData.foto +')' }"></div>
                        <div v-else-if="rinciaData.jenis == 'Perempuan'" class="avatar d-block rinciFoto" style="background-image: url(./assets/images/cewek.png)"></div>
                        <div v-else class="avatar d-block rinciFoto" style="background-image: url(./assets/images/cowok.png)"></div>
                        <div class="rincian-user">
                        <h5 style="text-transform:capitalize;">{{rinciaData.nama.toLowerCase()}}</h5>
                        <h6>Kelas {{rinciaData.kelas}}</h6>
                        </div>                        
                    </div>
                    <div v-if="rinciaData != ''" class="rincianPie">
                        <pie-charts :data="rinciaData"></pie-charts>
                    </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script type="application/javascript">
import moment from "moment";
import "moment/locale/id";
import PieCharts from "../../../components/PieCharts";
const today = new Date();
today.setHours(0, 0, 0, 0);

export default {
    components: {
        PieCharts,
    },
    data() {
        return {
            moment: moment,
            bulan: null,
            tableData: [],
            rinciaData:'',
            url: window.location.origin,
            pagination: [],            
            offset: 4,
            currentPage: 1,
            perPage: '20',
            search: '',
            sortedColumn: 'percent',
            order: 'asc'
        }
    },
    watch: {
      bulan(newVal) {
          if(newVal == null){        
            this.tableData=[];
            this.pagination=[];
            this.rinciaData="";
          }
      },
    },
    computed: {
        pagesNumber() {
            if (this.pagination.offset == undefined || this.pagination.offset == null) {
                return []
            }
            let from = this.pagination.page - this.offset
            if (from < 1) {
                from = 1
            }
            let to = from + (this.offset * 2)
            if (to >= this.pagination.total_page) {
                to = this.pagination.total_page
            }
            let pagesArray = []
            for (let page = from; page <= to; page++) {
                pagesArray.push(page)
            }
            return pagesArray
        },
        /**
         * Get the total data displayed in the current page.
         * */
        totalData() {
            return (this.pagination.offset - this.pagination.total_record) + 1
        }
    },
    methods: {
        /** Method **/
        fetchData() {
            let dataFetchUrl = `${this.url}/api/rekap/siswa?page=${this.currentPage}&tanggal=${this.bulan}&search=${this.search}&order=${this.sortedColumn}&by=${this.order}&limit=${this.perPage}`
            axios.get(dataFetchUrl)
                .then(response => {
                    this.pagination = response.data
                    this.tableData = response.data.records
                }).catch(error => this.tableData = [])
        },
        //limit tanggal
        notAfterToday(date) {
            return date > today;
        },
        // method percentage
        percentageData(data){
          var jumlah = data.hadir+data.izin+data.sakit+data.terlambat+data.alpa;
          var hadir = data.hadir+data.terlambat;
          return Math.floor(((hadir/jumlah)*100)+0.5);
        },
        onSort(){
          if(this.bulan != null){
            this.currentPage = 1;
            this.fetchData();
          }
        },
        onBulan() {
          if(this.bulan != null){
            this.fetchData();
          }
        },
        /** cari **/
        cari() {
          if(this.bulan != null){
            this.currentPage = 1;
            this.fetchData()
          }
        },
        /** enter search **/
        searchEnter: function(e) {
            if (e.keyCode === 13 && this.bulan != null) {
                this.currentPage = 1;
                this.fetchData() 
            }
        },
        /** refresh **/
        refresh() {
          if(this.bulan != null){
            this.search = "";
            this.fetchData();
          } 
        },
        changePage(pageNumber) {
          if(this.bulan != null){
            this.currentPage = pageNumber
            this.fetchData()
          }
        },
        rincian(data){
            this.rinciaData = '';
            this.rinciaData = data;            
        },
        //eend
    }
    //----
}
</script>