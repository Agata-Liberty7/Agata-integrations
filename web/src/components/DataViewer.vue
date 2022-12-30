<template>
  <div>
    <LineChart v-if="value.length" :chartData="chartData" :options="options" :plugins="plugins" />

    <button @click="page -= 1" :disabled="page < 1">
      назад
    </button>
    <button @click="page += 1">
      вперёд
    </button>
    <div v-if="value.length" style="margin: 10px; margin-left: 20px;">
      <vuetable ref="vuetable" :fields="fields" :api-mode="false" :data="value.slice(page * 20, (page + 1) * 20)"
        :sort-order="sortOrder" />
    </div>

    <h2 v-else style="margin-top: 10px;">Нет данных</h2>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import Vuetable from 'vue3-vuetable'
import { LineChart } from 'vue-chart-3';
import { Chart, registerables } from "chart.js";

Chart.register(...registerables);

const fieldNames = {
  'timestamp': 'Дата',
  'location': 'Локация',
  'name': 'Идентификатор',
  'value': 'Значение',
  'unit': 'Ед. изм.',
}

export default defineComponent({
  name: 'DataViewer',
  props: {
    value: Array,
    nameField: String,
    unitField: String,
  },
  components: {
    Vuetable,
    LineChart,
  },
  data: () => ({
    page: 0,
    sortOrder: [
      {
        field: 'year',
        direction: 'asc'
      }
    ]
  }),
  setup({ value, nameField = 'name' }) {
    const times = {}
    const names = {}

    value.forEach(item => {
      times[item.timestamp] = true
      names[item[nameField]] = true
    })
    const labels = Object.keys(times).sort()
    const units = Object.keys(names)

    const datasets = units.slice(0, 20).map((unit) => {

      const data = labels.map(label => {
        const item = value.find(item => item.timestamp === label && item[nameField] === unit)
        return item ? parseFloat(item.value) : NaN
      })

      return {
        label: unit,
        data,
        borderColor: '#' + Math.floor(Math.random() * 16777215).toString(16),
        cubicInterpolationMode: 'monotone',
        tension: 0.4,
        pointHoverRadius: 15,
        // radius: 0,
        fill: false,
        interaction: {
          intersect: false
        },
      }
    })

    const chartData = {
      labels: labels,
      datasets,
    };

    const plugins = [{
      decimation: {
        enabled: true,
        algorithm: 'lttb',
        samples: 50,
        // threshold: 0.5,
      },
    },
    ]

    const options = {
      responsive: true,
      scales: {
        x: {
          ticks: {
            callback: function (val, index, count) {
              if (count.length > 50) {
                return index % 5 === 0 ? this.getLabelForValue(val).slice(0, 7) : '';
              }
              return this.getLabelForValue(val).slice(0, 7);
            },
            color: 'grey',
          }
        },
      },
    };

    return { chartData, options, plugins };
  },
  computed: {
    fields: function () {
      if (this.value.length) {
        return Object.keys(this.value[0]).map(key => {
          return {
            name: key,
            title: fieldNames[key]
              ? `${fieldNames[key]} (${key})`
              : key,
            // sortField: key,
            // sortable: true,
            formatter(value) {
              if (key === 'timestamp') {
                const date = value.split(' ')[0]
                return date.split('-')[0] + "." + date.split('-')[1]
              }
              return value
            }
          }
        })
      }
      return []
    }
  },
})
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
