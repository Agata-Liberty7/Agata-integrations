<template>
  <div>
    <h1>Данные в системе</h1>

    <p>Выберите показатель для отображения данных, находящихся в системе. Если данных нет, нажмите <b>"Получить из
        реестра"</b> для загрузки данных из реестра в систему. Также возможно <b>"Сбросить данные"</b> по показателю.
      Чтобы вручную загрузить данные, <a href="/template.xlsx" target="_blank">скачайте шаблон</a>, заполните его и
      загрузите его в систему, нажав <b>"Выберите файл"</b>.
    </p>

    <select class="form-control" v-model="country" @change="changeCountry">
      <option value="-1" disabled>Выберите показатель</option>
      <option v-for="(v, k) in countries" :value="k" :key="k">{{ v }}</option>
    </select>

    <select class="form-control" v-model="section" @change="changeSection">
      <option value="-1" disabled>Выберите показатель</option>
      <option v-for="(v, k) in availableOptions" :value="k" :key="k">{{ v }}
      </option>
    </select>

    <button @click="fetchParam">
      Получить из реестра
    </button>
    <button @click="cleanParam">
      Сбросить данные
    </button>
    <input type="file" id="file" name="file" @change="uploadData" class="archive" />

    <div class="content">
      <div class="stats">
        <h2>Показатели</h2>
        <!-- <div class="item" v-for="s in stats" v-show="availableOptions.some(a => a.title === s.title)" :key="s.name" @click="() => changeSection({ target: { value: s.id } })"> -->
        <div class="item" v-for="s in stats" :key="s.name" @click="() => changeSection({ target: { value: s.name } })">
          <div>
            <span>{{ s.title }}:</span>
            <br />
            <span v-if="s.time" class="time">обновлено {{ s.time }}</span>
            <span v-else class="time-err">
              <red>ошибка парсинга</red>
            </span>
          </div>
          <b>{{ s.count }}</b>
        </div>
      </div>

      <div class="data">
        <h2 v-if="isLoading" style="margin-top: 10px;">Загрузка...</h2>
        <h2 v-if="isLoadingFed" style="margin-top: 10px;">Получение данных из реестра...</h2>
        <easy-spinner v-show="isLoading || isLoadingFed" color='#bada55' :size="30" />
        <data-viewer v-if="!isLoading && !isLoadingFed" :value="d"
          :nameField="(['population', 'vrp'].includes(section)) ? 'location' : 'name'" />
      </div>
    </div>
  </div>
</template>

<script setup>

</script>

<script>
import DataViewer from './DataViewer.vue'

const sections = {
  vvp: 'Индикаторы ВВП',
  currency: 'Курс валют',
  population: 'Население',
  retail: 'Количество ТТ',
  priceProducer: 'Цены производителей',
  priceConsumer: 'Потребительские цены',
  contracts: 'Контракты',
  priceAverage: 'Средняя цена по сегментам',
  vrp: 'Индикаторы ВРП',
}

const allCountries = {
  rus: {
    ru: 'Российская Федерация',
    be: 'Беларусь',
    kz: 'Казахстан',
    az: 'Азербайджан',
  },
  world: {
    kz: 'Казахстан',
    az: 'Азербайджан',
    cz: 'Чехия',
    de: 'Германия',
    us: 'США',
    eu: 'ЕС',
  }
}

export default {
  name: 'StatsPage',
  data: () => ({
    section: 'currency',
    country: 'ru',

    sections,
    availableOptions: { ...sections },

    isLoading: false,
    isLoadingFed: false,

    stats: [],
    d: [],
    isClients: false,
  }),
  components: {
    DataViewer,
  },
  computed: {
    countries: function () {
      return this.REGION === 'RUS'
        ? allCountries.rus
        : allCountries.world
    },
  },
  methods: {
    fetchParam() {
      this.isLoadingFed = true
      fetch(`${this.HOST}stats/${this.country}/${this.section}/fetch`)
        .then(response => {
          if (response.ok) {
            return response.json()
          }
          else {
            console.error("Server returned " + response.status + " : " + response.statusText);
          }
        })
        .then(response => {
          this.d = response
            ? response
              .sort((a, b) => parseInt(a.timestamp) - parseInt(b.timestamp))
              .sort((a, b) => ('' + a.location).localeCompare(b.location))
            : [];
          this.isLoadingFed = false
          this.getStats()
        })
    },

    cleanParam: function () {
      this.isLoading = true
      fetch(`${this.HOST}stats/${this.country}/${this.section}/clean`)
        .then(response => {
          if (response.ok) {
            return response.json()
          }
          else {
            console.error("Server returned " + response.status + " : " + response.statusText);
          }
        })
        .then(response => {
          this.d = response || []
          this.isLoading = false
          this.getStats()
        })
    },

    getData: function (country, section) {
      this.isLoading = true
      this.isLoadingFed = false
      fetch(`${this.HOST}stats/${country}/${section}`)
        .then(response => {
          if (response.ok) {
            return response.json()
          }
          else {
            console.error("Server returned " + response.status + " : " + response.statusText);
          }
        })
        .then(response => {
          if (response && response.length) {
            this.d = response
              .sort((a, b) => parseInt(b.timestamp) - parseInt(a.timestamp))
              .sort((a, b) => ('' + a.location).localeCompare(b.location))
          } else {
            this.d = []
          }
          this.isLoading = false
          this.getStats()
        })
    },


    uploadData: async function (e) {
      const url = `${this.HOST}stats/${this.country}/${this.section}/manual`

      var file = e.target.files[0];

      const formData = new FormData();
      console.log("file", file)
      formData.append('upload', file);
      formData.append('size', file.size);

      this.isLoading = true
      this.isLoadingFed = false

      fetch(url, {
        method: 'POST',
        body: formData
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
          if (response && response.length) {
            this.d = response
              .sort((a, b) => parseInt(b.timestamp) - parseInt(a.timestamp))
              .sort((a, b) => ('' + a.location).localeCompare(b.location))
          } else {
            this.d = []
          }
          this.isLoading = false
          this.getStats()
        })

      // this.request.action = 'upload';
      // this.request.file = files;
      // console.log("FILE " + this.request.file);

      // this.$http.post('data/getData.php', {
      //   data: this.request,
      //   headers: {
      //     'Content-Type': 'multipart/form-data'
      //   }
      // }
      // ).then(function (response) {

      //   console.log("RESPONSE: " + response.body);

      // }, function (response) {

      //   alert("error");
      // });

      // this.isLoading = true
      // this.isLoadingFed = false
      // fetch(`${this.HOST}stats/${country}/${section}/manual`)
      //   .then(response => {
      //     if (response.ok) {
      //       return response.json()
      //     }
      //     else {
      //       console.error("Server returned " + response.status + " : " + response.statusText);
      //     }
      //   })
      //   .then(response => {
      //     if (response && response.length) {
      //       this.d = response
      //         .sort((a, b) => parseInt(b.timestamp) - parseInt(a.timestamp))
      //         .sort((a, b) => ('' + a.location).localeCompare(b.location))
      //     } else {
      //       this.d = []
      //     }
      //     this.isLoading = false
      //     this.getStats()
      //   })

    },

    changeSection: function (e) {
      this.section = e.target.value
      this.$router.push({
        query: {
          ...this.$route.query,
          section: this.section,
        }
      });
      this.getData(this.country, this.section)
    },

    changeCountry: function (e) {
      this.country = e.target.value
      this.$router.push({
        query: {
          ...this.$route.query,
          country: this.country,
        }
      });
      this.getData(this.country, this.section)
    },

    getStats: function () {
      fetch(`${this.HOST}stats/${this.country}/count`)
        .then(response => {
          if (response.ok) {
            return response.json()
          }
          else {
            console.error("Server returned " + response.status + " : " + response.statusText);
          }
        })
        .then(response => {
          this.stats = [];
          this.availableOptions = {}

          for (let name of Object.keys(response)) {
            this.stats.push({
              name,
              title: this.sections[name],
              count: response[name].count,
              time: response[name].time,
            })

            this.availableOptions[name] = this.sections[name]
          }
        })
    },
  },
  created: function () {
    this.section = this.$route.query.section || 'currency';
    this.country = this.$route.query.country || (
      this.REGION === 'RUS' ? 'ru' : 'eu'
    );

    this.getData(this.country, this.section)
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
  padding: 8px;
  box-sizing: border-box;
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

.stats .item .time-err {
  font-size: .6rem;
  color: red;
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
