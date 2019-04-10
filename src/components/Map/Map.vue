<template>
  <div class="map-wrapper">
    <l-map
      style="height: inherit; width: 100%;"
      :zoom="zoom"
      :center="center"
      @update:zoom="zoomUpdated"
      @update:center="centerUpdated"
      @update:bounds="boundsUpdated"
      ref="map"
    >
      <l-tile-layer :url="url"></l-tile-layer>
      <map-circle
        v-for="(location, index) of cities"
        :key="index"
        :name="location.name"
        :engName="location.engName"
        :center="location.location"
        :color="'red'"
      />
    </l-map>
  </div>
</template>

<script>
import * as L from 'leaflet';
import {
  LMap, LTileLayer,
} from 'vue2-leaflet';
import MapCircle from './MapCircle.vue';
import locations from '../../data/locations';

export default {
  components: {
    LMap,
    LTileLayer,
    MapCircle,
  },
  data() {
    return {
      // url: "http://{s}.tile.stamen.com/watercolor/{z}/{x}/{y}.jpg",
      url: 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png',
      zoom: 4,
      center: [50.047, 20.0047],
      bounds: null,
    };
  },
  computed: {
    cities() {
      return locations;
    },
  },
  methods: {
    zoomUpdated(zoom) {
      this.zoom = zoom;
    },
    centerUpdated(center) {
      this.center = center;
    },
    boundsUpdated(bounds) {
      this.bounds = bounds;
    },
    showTripDetails(city) {
      const dest = locations.filter(c => c.name === city)[0];
      const group = L.latLngBounds([this.center, dest.location]);
      this.$refs.map.fitBounds(group);
    },
    updateMap() {
      setTimeout(() => this.$refs.map.mapObject.invalidateSize(), 1);
    },
  },
};
</script>

<style lang="scss">
.map-wrapper {
  width: 100%;
  height: 100vh;
}
</style>
