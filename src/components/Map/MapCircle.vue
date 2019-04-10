<template>
  <l-feature-group v-if="hits > 0">
    <l-marker :lat-lng="center">
      <l-tooltip>
        <b>{{name}}</b>
        <br>
        <span style="color: gray;">Liczba osób: {{hits}}</span>
      </l-tooltip>
    </l-marker>
    <l-circle :lat-lng="center" :radius="radius" :color="color" :fill-color="color"></l-circle>
    <l-circle
      v-if="hitsDone > 0"
      :lat-lng="center"
      :radius="(hitsDone / hits) * radius"
      :color="'green'"
      :fill-color="'green'"
    ></l-circle>
  </l-feature-group>
</template>

<script>
import {
  LCircle, LTooltip, LFeatureGroup, LMarker,
} from 'vue2-leaflet';
import { getRequestsCollection } from '../../firebase';

export default {
  components: {
    LCircle,
    LTooltip,
    LMarker,
    LFeatureGroup,
  },
  props: ['center', 'name', 'engName'],
  data() {
    return {
      color: '#0C81E4',
      hits: 0,
      radius: 0,
      hitsDone: 0,
      maxRadius: 400000,
      minRadius: 80000,
      radiusIncrement: 10000,
    };
  },
  mounted() {
    getRequestsCollection().where('canceled', '==', false).onSnapshot(((snapshot) => {
      const totalHits = snapshot.docs.length;
      const requestsForCity = snapshot.docs.filter(r => r.data().city === this.engName);
      this.hits = requestsForCity.length;
      this.hitsDone = requestsForCity.filter(r => r.data().status === 'done').length;
      this.radius = Math.min(this.hits * this.radiusIncrement + this.minRadius, this.maxRadius);
    }));
  },
};
</script>

<style>
.map-circle {
  transition-property: stroke, fill;
  transition-duration: 0.5s;
  transition-timing-function: ease-in-out;
}
</style>
