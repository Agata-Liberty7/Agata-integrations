<template>
  <div class="hello">

    <div class="tabs">
      <div :class="'tab ' + (tab === 0 ? 'active' : '')" @click="() => changeTab(0)">Реестры</div>
      <div :class="'tab ' + (tab === 1 ? 'active' : '')" @click="() => changeTab(1)">Данные клиентов</div>
      <div :class="'tab ' + (tab === 2 ? 'active' : '')" @click="() => changeTab(2)">Настройки</div>
    </div>

    <StatsPage v-if="!tab" />
    <ClientsPage v-if="tab === 1" />
    <MetadataPage v-if="tab === 2" />
  </div>
</template>

<script>
import MetadataPage from './Metadata.vue'
import ClientsPage from './Clients.vue'
import StatsPage from './Stats.vue'

// const host = 'http://80.249.144.41:80/'
// const host = 'http://178.62.240.212:8080/'
// const host = 'http://localhost:8080/'

export default {
  name: 'HelloWorld',
  components: {
    MetadataPage,
    ClientsPage,
    StatsPage,
  },
  data: () => ({
    tab: 0,
  }),
  methods: {
    changeTab: function(tab) {
      this.tab = tab
      console.log(this.msg)
    },
  },
}
</script>

<style scoped>
 .tabs {
    display: flex;
    align-items: center;
 }

 .tab {
    cursor: pointer;
    padding: 5px 10px;
    border-radius: 5px;
    border: 1px solid #ccc;
    color: #999;
    font-size: 14px;
    font-weight: bold;
    transition: all 0.2s;
 }

 .tab.active {
    background: #fff;
    color: #333;
    text-transform: uppercase;
 }
</style>
