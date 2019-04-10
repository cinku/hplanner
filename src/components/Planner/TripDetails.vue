<template>
  <div class="details-wrapper" v-if="tripData">
    <div class="details__destination">
      <div>{{tripData.city}}</div>
      <div>
        <span class="details__time">
          {{tripTime(tripData.flightDate, tripData.returnDate)}} DNI
        </span>
        <span class="details__cost">
          {{tripCost(tripData.flightCost, tripData.hotelCost)}}
        </span>
      </div>
    </div>
    <div class="trip-card">
      <div class="trip-card__header">
        <div class="trip-card__title">
          <i class="trip-icon fas fa-plane-departure"></i>
          Lot
        </div>
        <div class="trip-card__info">{{flightSeparateCost(tripData.flightCost)}}</div>
      </div>
      <div class="flight-date">{{flightDay(tripData.flightDate)}}</div>
      <div class="trip-card__content">
        <div class="flight-details">
          <div>
            <div class="flight-label">WYLOT</div>
            <div class="flight-time">{{flightDepartureTime(tripData.flightDate)}}</div>
          </div>
          <div class="flight-details--icon">
            <i class="fas fa-long-arrow-alt-right"></i>
          </div>
          <div>
            <div class="flight-label">PRZYLOT</div>
            <div class="flight-time">{{flightArrivalTime(tripData.flightDate)}}</div>
          </div>
        </div>
      </div>
    </div>
    <div class="trip-card">
      <div class="trip-card__header">
        <div class="trip-card__title">
          <i class="trip-icon fas fa-home"></i>
          Hotel
        </div>
        <div class="trip-card__info">{{hotelCost(tripData.hotelCost)}}</div>
      </div>
      <div class="trip-card__content">
        <a :href="tripData.hotelUrl" target="_blank">{{tripData.hotelName}}</a>
      </div>
    </div>
    <div class="trip-card">
      <div class="trip-card__header">
        <div class="trip-card__title">
          <i class="trip-icon fas fa-plane-arrival"></i>
          Lot
        </div>
        <div class="trip-card__info">{{flightSeparateCost(tripData.flightCost)}}</div>
      </div>
      <div class="flight-date">{{flightDay(tripData.returnDate)}}</div>
      <div class="trip-card__content">
        <div class="flight-details">
          <div>
            <div class="flight-label">WYLOT</div>
            <div class="flight-time">{{flightDepartureTime(tripData.returnDate)}}</div>
          </div>
          <div class="flight-details--icon">
            <i class="fas fa-long-arrow-alt-right"></i>
          </div>
          <div>
            <div class="flight-label">PRZYLOT</div>
            <div class="flight-time">{{flightArrivalTime(tripData.returnDate)}}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import moment from 'moment';

export default {
  props: ['tripData'],
  methods: {
    flightDay(date) {
      return date.split(' ')[0];
    },
    flightDepartureTime(date) {
      return date.split(' ')[1];
    },
    flightArrivalTime(date) {
      return date.split(' ')[2];
    },
    flightSeparateCost(cost) {
      let separateCost = cost.replace('€', '');
      separateCost = Math.floor(+separateCost / 2);
      return `${separateCost}€`;
    },
    hotelCost(cost) {
      const filteredCost = cost.replace('€', '').trim();
      return `${filteredCost}€`;
    },
    tripTime(startDate, endDate) {
      const start = moment(startDate.split(' ')[0], 'DD-MM-YY');
      const end = moment(endDate.split(' ')[0], 'DD-MM-YY');
      return end.diff(start, 'days');
    },
    tripCost(flightCost, hotelCost) {
      const flight = flightCost.replace('€', '');
      const hotel = hotelCost.replace('€', '');
      const cost = +hotel + +(Math.floor(flight / 2) * 2);
      return `${cost}€`;
    },
  },
};
</script>

<style lang="scss" scoped>
.details-wrapper {
  padding: 5px;
  text-align: left;
}

.details__destination {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 15px;
  color: #6f7386;
  display: flex;
  justify-content: space-between;

  &-icon {
    color: #55b0f3;
    margin-left: 10px;
    font-size: 20px;
  }
}

.details__time {
  color: #6f7386;
  font-size: 14px;
  font-weight: normal;
}

.details__cost {
  color: #55b0f3;
  margin-left: 10px;
}

.trip-card {
  background: #f2f4f8;
  padding: 10px 25px;
  text-align: left;
  color: #333;
  border-radius: 5px;
  margin-bottom: 15px;
  font-size: 20px;
}

.trip-card__header {
  display: flex;
  justify-content: space-between;
}

.trip-card__info {
  color: #55b0f3;
  font-weight: 600;
}

.trip-card__content {
  margin-top: 10px;
  font-weight: 600;
  & > a {
    color: #55b0f3;
  }
}

.trip-icon {
  color: #fca97d;
  margin-right: 10px;
}

.flight-details {
  display: flex;
  justify-content: space-between;

  &--icon {
    font-size: 30px;
    color: #55b0f3;
  }
}

.flight-label {
  color: #bdbdbd;
  font-size: 12px;
  font-weight: normal;
  margin-bottom: -5px;
}

.flight-time {
  color: #6f7386;
}

.flight-date {
  color: #6f7386;
  font-size: 14px;
  margin-top: 10px;
}
</style>
