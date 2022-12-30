<template>
  <div>
    <h1>Данные клиентов</h1>
    <p>Выберите клиента для отображения данных, находящихся в системе. Если данных нет, нажмите <b>"Получить из
        реестра"</b> для загрузки данных из реестра в систему. Также возможно <b>"Сбросить данные"</b> по показателю.
    </p>


    <select v-model="clientId" @change="() => currentClientData = []">
      <option v-for="c in clients" v-bind:value="c.id" v-bind:key="c.id">{{ c.name }}</option>
    </select>

    <div class="content">
      <div class="addFields">
        <h4>Добавить данные</h4>
        <div v-for="(f, i) in currentClientData" class="field" v-bind:key="i">
          <select @change="$event => f.map = $event.target.value">
            <option v-for="c in fieldTypes" v-bind:value="c.map" v-bind:key="c.id">{{ c.name }}</option>
          </select>
          <input type="text" v-model="f.value" />
          <button
            @click="() => currentClientData = currentClientData.filter((item) => !(item.map === f.map && item.value === f.value))">-</button>
        </div>

        <button @click="() => currentClientData = [...currentClientData, { map: 'sales', value: '' }]">+</button>
        <button v-if="currentClientData.length" @click="postClient()">Сохранить</button>
      </div>

      <div class="data">
        <!-- <div v-for="client in clientData" v-bind:key="client.clientId"> -->
        <h4>История операций</h4>
        <div v-if="clientData[clientId]">
          <p v-for="(d, i) in clientData[clientId].items" v-bind:key="i">
            <span>{{ d.timestamp }} </span>
            <span style="color: green; font-weight: bold;">{{ d.status }}</span>
          </p>
        </div>

        <h4>Текущие данные клиента</h4>
        <div v-if="clientData[clientId]">
          <p v-for="item in formatClientData(clientData[clientId].data || {})" :key="item.name">
            <span>{{ item.name }}: </span>
            <b>{{ item.value }}</b>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

const fieldTypes = [
  { id: 0, name: "Продукт", map: "product" },
  { id: 1, name: "Ед.изм.", map: "unit" },
  { id: 2, name: "Вес 1 шт, кг	", map: "unit_weight" },
  { id: 3, name: "Кол-во ед.хранения", map: "count" },
  { id: 4, name: "Кол-во, кг	", map: "weight" },
  { id: 5, name: "Продажи Руб без НДС", map: "sales" },
  { id: 6, name: "Продажи Руб с НДС", map: "sales_nds" },
  { id: 7, name: "Канал дистрибуции", map: "distribution" },
  { id: 8, name: "Регион отгрузки", map: "region" },
  { id: 9, name: "Контрагент", map: "counterparty" },
  { id: 10, name: "Бренд", map: "brand" },
  { id: 11, name: "Товарная категория", map: "product_category" },
  { id: 12, name: "ФО/Экспорт", map: "export" },
  { id: 13, name: "Год", map: "year" },
  { id: 14, name: "Период", map: "period" },
]

const clients = [
  { id: 0, name: "Клиент #1" },
  { id: 1, name: "Клиент #2" },
  { id: 2, name: "Клиент #3" },
]


export default {
  name: 'ClientsPage',
  data: () => ({
    data: [],
    clients,
    fieldTypes,
    clientId: 0,
    currentClientData: [
      // { map: "", name: "", value: ""}
    ],
    clientData: [],
  }),
  methods: {
    formatClientData: function (data) {
      // const result = []
      // data.forEach((item) => {
      //   const timestamp = new Date(item.timestamp)
      //   const date = `${timestamp.getDate()}.${timestamp.getMonth() + 1}.${timestamp.getFullYear()}`
      //   const time = `${timestamp.getHours()}:${timestamp.getMinutes()}`
      //   const status = item.status
      //   const map = item.map
      //   const value = item.value
      //   result.push({ timestamp: `${date} ${time}`, status, map, value })
      // })
      // return result

      // const result = []

      return Object.keys(data).map(key => ({
        value: data[key],
        name: fieldTypes.find(item => item.map === key).name,
      })
      )
    },

    postClient: function () {
      this.isLoading = true
      fetch(`${this.HOST}clients/${this.clientId}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          fields: this.currentClientData,
        })
      })
        .then(response => {
          if (response.ok) {
            return response.json()
          }
          else {
            console.error("Server returned " + response.status + " : " + response.statusText);
          }
        })
        .then(response => {
          this.clientData = response
            ? response
            // .sort((a, b) => parseInt(b.timestamp) - parseInt(a.timestamp))
            // .sort((a, b) => ('' + a.location).localeCompare(b.location)) 
            : [];
          this.isLoading = false
          this.fetchClients()
          this.currentClientData = []
        })
    },

    created: function () {
      this.fetchClients()
    },
  },
}
</script>


<style scoped>
.content {
  display: flex;
  flex-direction: row;
}

.stats {
  width: 200px;
  text-align: justify;
  margin-top: -3px;
  font-size: .8rem;
}

.stats .item {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 10px;
  cursor: pointer;
}

.stats .item .time {
  font-size: .6rem;
  color: #a0a0a0;
  margin-top: -3px;
  display: block;
}

.data,
.addFields {
  flex: 1;
}

.field {
  display: flex;
}

@media screen and (max-width: 768px) {
  .content {
    flex-direction: column;
  }

  .stats {
    width: 100%;
  }

  button {
    margin-top: 10px;
  }
}

h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

button {
  margin-left: 10px;
  padding: 5px 10px;
}

select,
input {
  padding: 5px 10px;
  margin: 5px;
}
</style>
