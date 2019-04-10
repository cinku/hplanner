<template>
  <l-feature-group v-if="hits > 0">
    <l-marker :lat-lng="center">
      <l-tooltip>
        <b>{{name}}</b><br />
        <span v-if="hits" style="color: gray;">
          Liczba osób: {{hits}}
        </span>
      </l-tooltip>
    </l-marker>
    <l-circle :lat-lng="center" :radius="radius" :color="color" :fill-color="color"></l-circle>
    <l-circle
      v-if="doneRatio > 0"
      :lat-lng="center"
      :radius="doneRatio * radius"
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
      maxRadius: 400000,
      minRadius: 80000,
      doneRatio: 0,
      growthSlowingFactor: 10,
    };
  },
  mounted() {
    getRequestsCollection().where('canceled', '==', false).onSnapshot(((snapshot) => {
      const totalHits = snapshot.docs.length;
      if (totalHits > 0) {
        const requestsForCity = snapshot.docs.filter(r => r.data().city === this.engName);
        this.hits = requestsForCity.length;
        this.hitsDone = requestsForCity.filter(r => r.data().status === 'done').length;
        this.doneRatio = this.hitsDone / this.hits;
        const radiusRange = this.maxRadius - this.minRadius;
        this.radius = (this.hits / totalHits) * radiusRange
          * (Math.min(this.growthSlowingFactor, totalHits) / this.growthSlowingFactor)
          + this.minRadius;
      } else {
        this.hits = 0;
        this.hitsDone = 0;
        this.doneRatio = 0;
        this.radius = 0;
      }
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
