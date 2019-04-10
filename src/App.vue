<template>
  <div id="app">
    <div class="planner-section" :class="{ 'map-section--hidden': !mapHidden, 'map-section--shown': mapHidden }">
      <Planner :show-trip="showTrip" :trip-data="tripData" :searching="searching" @submitLocation="handleSubmit" @cancel="cancel" />
    </div>
    <div class="map-section" :class="{ 'map-section--hidden': mapHidden, 'map-section--shown': !mapHidden }">
      <Map ref="map"/>
    </div>
    <button class="toggle-map-button" @click="toggleMap" title="Toggle map"><i class="fas fa-map"></i></button>
  </div>
</template>

<script>
import Map from './components/Map/Map.vue';
import Planner from './components/Planner/Planner.vue';
import { getRequestsCollection, getRequestById } from './firebase';

export default {
  name: 'app',
  components: {
    Map,
    Planner,
  },
  data() {
    return {
      searching: false,
      showTrip: false,
      tripData: null,
      mapHidden: true,
    };
  },
  methods: {
    toggleMap() {
      this.mapHidden = !this.mapHidden;
      this.$refs.map.updateMap();
    },
    handleSubmit({
      city, month, name, email,
    }) {
      this.searching = true;
      this.showTrip = false;
      getRequestsCollection()
        .add({
          name,
          email,
          city,
          month,
          canceled: false,
          status: 'processing',
        })
        .then((docRef) => {
          localStorage.referenceId = docRef.id;
          getRequestById(localStorage.referenceId).onSnapshot((doc) => {
            if (doc.exists) {
              const data = doc.data();
              if (data.status === 'processing') {
                this.searching = true;
              } else if (data.status === 'done') {
                this.searching = false;
                this.tripData = data;
                this.showTrip = true;
                this.$refs.map.showTripDetails(this.tripData.city);
              }
            }
          });
        });
    },
    cancel() {
      getRequestById(localStorage.referenceId).update({ canceled: true })
        .then(() => {
          this.showTrip = false;
          this.searching = false;
          this.tripData = null;
          localStorage.removeItem('referenceId');
        });
    },
  },
  created() {
    if (localStorage.referenceId) {
      this.showTrip = true;
      getRequestById(localStorage.referenceId).onSnapshot((doc) => {
        if (doc.exists) {
          const data = doc.data();
          if (data.status === 'processing') {
            this.searching = true;
          } else if (data.status === 'done') {
            this.searching = false;
            this.tripData = data;
            this.showTrip = true;
            this.$refs.map.showTripDetails(this.tripData.city);
          }
        } else {
          localStorage.removeItem('referenceId');
        }
      });
    }
  },
};
</script>

<style lang="scss">
#app {
  font-family: 'Open Sans', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  padding: 0;
  display: flex;
  position: relative;
}

body {
  margin: 0;
}

.map-section {
  width: calc(100vw - 400px);
  z-index: 2;

  @media screen and (max-width: 1024px) {
    width: 100%;

    &--hidden {
      display: none;
    }

    &--shown {
      display: block;
      z-index: 15;
    }
  }
}
.planner-section {
  width: 400px;
  min-height: 100vh;
  box-shadow: 3px 0px 16px rgba(0, 0, 0, 0.2);
  z-index: 3;
  @media screen and (max-width: 1024px) {
    width: 100%;
  }
}

.toggle-map-button {
  display: none;
  @media screen and (max-width: 1024px) {
    display: block;
    position: fixed;
    z-index: 100;
    bottom: 15px;
    right: 15px;
    background: #55B0F3;
    border-radius: 50%;
    color: #fff;
    cursor: pointer;
    font-size: 19px;
    padding:0;
    width: 50px;
    height: 50px;
    box-shadow: 0px 2px 12px rgba(0, 0, 0, 0.31);
    transition: background 0.3s ease-in-out;
    border: none;
    &:hover {
      background: darken(#55B0F3, 10%);
    }
  }
}

.form-wrapper {
  text-align: left;
  margin-top: 20px;
  @media screen and (max-width: 1024px) {
    margin-top: 0;
  }
}

.form-group {
  @media screen and (max-width: 1024px) {
    margin-bottom: 5px;
  }
  & > label {
    margin-bottom: 0;
    font-size: 14px;
    font-weight: 600;
  }

  .input-group-text {
    border: none;
    color: #c5c5c5;
    background: #f5f4f8 !important;
  }
}

input,
select {
  background: #f5f4f8 !important;
  border-radius: 5px !important;
  border: none !important;
  font-size: 14px !important;

  &::placeholder {
    font-size: 14px;
  }
}

.submit-button {
  border-radius: 30px;
  margin-top: 40px;
  border-width: 3px;
  font-weight: bold;

  &:disabled {
    cursor: not-allowed;
  }

  @media screen and (max-width: 1024px) {
    margin-top: 20px;
  }
}

.search-icon {
  margin-left: 10px;
}
</style>
