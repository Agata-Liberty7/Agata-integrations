<template>
  <div>
    <!-- <button @click="fetchMeta">load</button> -->
    <vuetable ref="vuetable2" :fields="fields" :data="data" :api-mode="false">
      <template #vvp="props">
        <select :disabled="props.rowData.vvp === undefined" @change="changeMeta(props.rowData.location, 'vvp', $event)"
          v-model="props.rowData.vvp">
          <option value="day">Каждый день</option>
          <option value="week">Каждую неделю</option>
          <option value="month">Каждый месяц</option>
          <option value="quartal">Каждый квартал</option>
          <option value="year">Каждый год</option>
        </select>
      </template>

      <template #vrp="props">
        <select :disabled="props.rowData.vrp === undefined" @change="changeMeta(props.rowData.location, 'vrp', $event)"
          v-model="props.rowData.vrp">
          <option value="day">Каждый день</option>
          <option value="week">Каждую неделю</option>
          <option value="month">Каждый месяц</option>
          <option value="quartal">Каждый квартал</option>
          <option value="year">Каждый год</option>
        </select>
      </template>

      <template #currency="props">
        <select :disabled="props.rowData.currency === undefined"
          @change="$event => changeMeta(props.rowData.location, 'currency', $event)" v-model="props.rowData.currency">
          <option value="day">Каждый день</option>
          <option value="week">Каждую неделю</option>
          <option value="month">Каждый месяц</option>
          <option value="quartal">Каждый квартал</option>
          <option value="year">Каждый год</option>
        </select>
      </template>

      <template #population="props">
        <select :disabled="props.rowData.population === undefined"
          @change="$event => changeMeta(props.rowData.location, 'population', $event)"
          v-model="props.rowData.population">
          <option value="day">Каждый день</option>
          <option value="week">Каждую неделю</option>
          <option value="month">Каждый месяц</option>
          <option value="quartal">Каждый квартал</option>
          <option value="year">Каждый год</option>
        </select>
      </template>

      <template #retail="props">
        <select :disabled="props.rowData.retail === undefined"
          @change="$event => changeMeta(props.rowData.location, 'retail', $event)" v-model="props.rowData.retail">
          <option value="day">Каждый день</option>
          <option value="week">Каждую неделю</option>
          <option value="month">Каждый месяц</option>
          <option value="quartal">Каждый квартал</option>
          <option value="year">Каждый год</option>
        </select>
      </template>

      <template #priceProducer="props">
        <select :disabled="props.rowData.priceProducer === undefined"
          @change="$event => changeMeta(props.rowData.location, 'priceProducer', $event)"
          v-model="props.rowData.priceProducer">
          <option value="day">Каждый день</option>
          <option value="week">Каждую неделю</option>
          <option value="month">Каждый месяц</option>
          <option value="quartal">Каждый квартал</option>
          <option value="year">Каждый год</option>
        </select>
      </template>

      <template #priceConsumer="props">
        <select :disabled="props.rowData.priceConsumer === undefined"
          @change="$event => changeMeta(props.rowData.location, 'priceConsumer', $event)"
          v-model="props.rowData.priceConsumer">
          <option value="day">Каждый день</option>
          <option value="week">Каждую неделю</option>
          <option value="month">Каждый месяц</option>
          <option value="quartal">Каждый квартал</option>
          <option value="year">Каждый год</option>
        </select>
      </template>

      <template #contracts="props">
        <select :disabled="props.rowData.contracts === undefined"
          @change="$event => changeMeta(props.rowData.location, 'contracts', $event)" v-model="props.rowData.contracts">
          <option value="day">Каждый день</option>
          <option value="week">Каждую неделю</option>
          <option value="month">Каждый месяц</option>
          <option value="quartal">Каждый квартал</option>
          <option value="year">Каждый год</option>
        </select>
      </template>

      <template #priceAverage="props">
        <select :disabled="props.rowData.priceAverage === undefined"
          @change="$event => changeMeta(props.rowData.location, 'priceAverage', $event)"
          v-model="props.rowData.priceAverage">
          <option value="day">Каждый день</option>
          <option value="week">Каждую неделю</option>
          <option value="month">Каждый месяц</option>
          <option value="quartal">Каждый квартал</option>
          <option value="year">Каждый год</option>
        </select>
      </template>
    </vuetable>

  </div>
</template>

  <script>
import Vuetable from 'vue3-vuetable'

// const tmpData = [
//   {
//     location: 'РОССТАТ',
//     vvp: 'month',
//     vrp: 'day',
//     currency: 'year',
//     population: 'year',
//     retail: 'year',
//     priceProducer: 'year',
//     priceConsumer: 'year',
//     contracts: 'year',
//     priceAverage: 'year',
//   },
//   {
//     location: 'КАЗСТАТ',
//     currency: 'year',
//     population: 'year',
//   },
//   {
//     location: 'УКРСТАТ',
//     currency: 'year',
//     population: 'year',
//   },
//   {
//     location: 'БЕЛСТАТ',
//     currency: 'year',
//     population: 'year',
//   },
//   {
//     location: 'ЕВРОСТАТ',
//     vrp: 'day',
//     currency: 'year',
//     population: 'year',
//   },
//   {
//     location: 'США',
//   },
//   {
//     location: 'Германия',
//   },
//   {
//     location: 'Великобритания',
//   },
// ]

// const fieldNames = {
//   'timestamp': 'Дата',
//   'location': 'Локация',
//   'name': 'Идентификатор',
//   'value': 'Значение',
//   'unit': 'Ед. изм.',
// }

const options = [
  { id: -1, name: 'location', title: 'Реестр' },
  { id: 0, name: 'vvp', title: 'Индикаторы ВВП' },
  { id: 1, name: 'currency', title: 'Курс валют' },
  { id: 2, name: 'population', title: 'Население' },
  { id: 3, name: 'retail', title: 'Количество ТТ' },
  { id: 4, name: 'priceProducer', title: 'Цены производителей' },
  { id: 5, name: 'priceConsumer', title: 'Потребительские цены' },
  { id: 6, name: 'contracts', title: 'Контракты' },
  { id: 7, name: 'priceAverage', title: 'Средняя цена по сегментам' },
  { id: 8, name: 'vrp', title: 'Индикаторы ВРП' },
]

export default {
  name: 'MetadataPage',
  // props: {
  //   data: Array
  // },
  components: {
    Vuetable,
  },
  data: () => ({
    // tmpData,
    // data: [...tmpData],
    sortOrder: [
      {
        field: 'year',
        direction: 'asc'
      }
    ]
  }),
  computed: {
    data: function () {
      return this.REGION === 'RUS'
        ? [
          {
            location: 'РОССТАТ',
            vvp: 'month',
            vrp: 'day',
            currency: 'year',
            population: 'year',
            retail: 'year',
            priceProducer: 'year',
            priceConsumer: 'year',
            contracts: 'year',
            priceAverage: 'year',
          },
          {
            location: 'КАЗСТАТ',
            population: 'year',
            vrp: 'year',
            retail: 'year',
          },
          {
            location: 'БЕЛСТАТ',
            population: 'year',
            vrp: 'year',
            retail: 'year',
          },
          {
            location: 'Азербайджан',
            population: 'year',
            vrp: 'year',
            retail: 'year',
          },
        ]
        : [
          {
            location: 'Казахстан',
            population: 'year',
            vrp: 'year',
            retail: 'year',
          },
          {
            location: 'Азербайджан',
            population: 'year',
            vrp: 'year',
            retail: 'year',
          },
          {
            location: 'Чехия',
            vrp: 'year',
            population: 'year',
          },
          {
            location: 'ЕС',
            vrp: 'day',
            currency: 'year',
            population: 'year',
          },
          {
            location: 'США',
            population: 'year',
            vrp: 'year',
          },
          {
            location: 'Германия',
            population: 'year',
            vrp: 'year',
            retail: 'year',
          },
        ]
    },
    // fields: function () {
    //   if (this.data.length) {
    //     return Object.keys(this.data[0]).map(key => {
    //       return {
    //         name: key,
    //         title: fieldNames[key] || key,
    //         // sortField: key,
    //         // sortable: true,
    //         formatter (value) {
    //           if (key === 'timestamp') {
    //             const date = value.split(' ')[0]
    //             return date.split('-')[0] + "." + date.split('-')[1]
    //           }
    //           return value
    //         }
    //       }
    //     })
    //   }
    //   return []
    // }

    fields: function () {
      // if (this.data.length) {
      return options.map(option => {
        return {
          name: option.name,
          title: option.title,
          // sortField: key,
          // sortable: true,
          // formatter (value) {
          //   if (key === 'timestamp') {
          //     const date = value.split(' ')[0]
          //     return date.split('-')[0] + "." + date.split('-')[1]
          //   }
          //   return value
          // }
        }
      })
      // }
      // return []
    },

    resData: function () {
      // let res = []

      // for (let i = 0; i < this.data.length; i++) {
      //   let item = this.data[i]
      //   let itemRes = {}
      //   for (let j = 0; j < options.length; j++) {
      //     let option = options[j]
      //     itemRes[option.name] = item[option.name]
      //   }
      //   res.push(itemRes)
      // }

      // Object.keys(this.data[0]).map(key => {

      // })


      return this.data.map(item => {
        return {
          location: item.location,
          vvp: item.vvp,
          currency: item.currency,
          population: item.population,
          retail: item.retail,
          priceProducer: item.priceProducer,
          priceConsumer: item.priceConsumer,
          contracts: item.contracts,
          priceAverage: item.priceAverage,
          vrp: item.vrp,
        }
      })
    }
  },
  methods: {

    fetchMeta: function () {
      // this.isLoadingFed = true
      fetch(`${this.HOST}meta`)
        .then(response => {
          if (response.ok) {
            return response.json()
          }
          else {
            console.error("Server returned " + response.status + " : " + response.statusText);
          }
        })
        .then(response => {
          // this.data = response
          // ? response
          // : []; 
          console.log(response)
          // this.isLoadingFed = false
          // this.getStats()
        })
    },

    changeMeta: function (location, section, value) {
      console.log(location, section, value.target.value)
      const index = this.data.indexOf(d => d.location === location)
      console.log(index)
      console.log(this.data)
      // this.data[index][section] = value.target.value
    },
  },

  created: function () {
    this.fetchMeta()
  },
}
</script>

  <!-- Add "scoped" attribute to limit CSS to this component only -->
  <style scoped>
  h3 {
    margin: 40px 0 0;
  }
  
  li {
    margin: 0 10px;
  }
  
  a {
    color: #42b983;
  }
  </style>
