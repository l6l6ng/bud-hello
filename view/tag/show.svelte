<script>
  export let tag = {}
</script>

<h1>Showing Tag</h1>

<table border="1" cellpadding="10">
  <thead>
    {#each Object.keys(tag) as key}
      <th>{key}</th>
    {/each}
  </thead>
  <tr>
    {#each Object.keys(tag) as key}
      <td>{tag[key]}</td>
    {/each}
  </tr>
</table>

<br />

<a href="..">Back</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
