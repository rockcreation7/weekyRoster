<template>
  <div class="submit-form">
    <div v-if="!submitted">
      <div class="form-group">
        <label for="date">Date</label>
        <input
          type="text"
          class="form-control"
          id="date"
          required
          v-model="roster.date"
          name="date"
          list="dateList"
        />
      </div>

      <div class="form-group">
        <label for="upperstaff">Upper Staff</label>
        <input
          class="form-control"
          id="upperstaff"
          required
          v-model="roster.upperstaff"
          name="upperstaff"
          list="staff"
        />
      </div>

      <div class="form-group">
        <label for="uppertime">Upper Time</label>
        <input
          class="form-control"
          id="uppertime"
          required
          v-model="roster.uppertime"
          name="uppertime"
          list="time"
        />
      </div>
      <div class="form-group">
        <label for="lowerstaff">Lower Staff</label>
        <input
          class="form-control"
          id="lowerstaff"
          required
          v-model="roster.lowerstaff"
          name="lowerstaff"
          list="staff"
        />
      </div>

      <div class="form-group">
        <label for="lowertime">Lower Time</label>
        <input
          class="form-control"
          id="lowertime"
          required
          v-model="roster.lowertime"
          name="lowertime"
          list="time"
        />
      </div>

      <div class="form-group">
        <label for="custommessage">Custom message</label>
        <input
          class="form-control"
          id="custommessage"
          required
          v-model="roster.custommessage"
          name="custommessage"
        />
      </div> 
      <DataList/>
      <button @click="saveRoster" class="btn btn-success">Submit</button>
    </div>
    <div v-else>
      <h4>You submitted successfully!</h4>
      <button class="btn btn-success" @click="newRoster">Add Roster</button>
    </div>
  </div>
</template>

<script>
import RosterDataService from "../services/RosterDataService";
import DataList from "../components/DataList";
export default {
  components: {
    DataList
  },
  name: "add-roster",
  data() {
    return {
      roster: {
        date: "",
        upperstaff: "",
        uppertime: "",
        lowerstaff: "",
        lowertime: "",
        custommessage: "",
      },
      submitted: false,
    };
  },
  methods: {
    saveRoster() {
      var data = {
        date: this.roster.date,
        upperstaff: this.roster.upperstaff,
        uppertime: this.roster.uppertime,
        lowerstaff: this.roster.lowerstaff,
        lowertime: this.roster.lowertime,
      };

      RosterDataService.create(data)
        .then((response) => {
          this.submitted = true;
          console.log(response);
        })
        .catch((e) => {
          console.log(e);
        });
    },

    newRoster() {
      this.submitted = false;
      this.tutorial = {};
    },
  },
};
</script>

<style>
.submit-form {
  max-width: 300px;
  margin: auto;
}
</style>