<template>
  <div>

    <datalist id="staff">
      <option value="Chloe"></option>
      <option value="Fanny"></option>
      <option value="Nam"></option>
      <option value="Nicole"></option>
      <option value="Yo"></option>
    </datalist>

    <datalist id="time">
      <option value="11:00 am - 04:00 pm"></option>
      <option value="11:00 am - 05:00 pm"></option>
      <option value="11:00 am - 07:00 pm"></option>
      <option value="11:00 am - 08:00 pm"></option>
      <option value="12:00 am - 06:00 pm"></option>
      <option value="12:00 am - 08:00 pm"></option>
      <option value="6:00 pm - 08:00 pm"></option>
    </datalist>

    <datalist id="dateList">
      <option :value="new Date().getFullYear().toString()"></option>
    </datalist>
  </div>
</template>
 
 