<script>
  export let tags = []
</script>

<h1>Tag Index</h1>

<table border="1" cellpadding="10">
  {#if tags.length > 0}
    <thead>
      {#each Object.keys(tags[0]) as key}
        <th>{key}</th>
      {/each}
    </thead>
  {/if}
  {#each tags as tag}
    <tr>
      {#each Object.keys(tag) as key}
        {#if key === "id"}
          <td><a href={`/${tag[key]}`}>{tag[key]}</a></td>
        {:else}
          <td>{tag[key]}</td>
        {/if}
      {/each}
    </tr>
  {/each}
</table>

<style>
  table {
    border-collapse: collapse;
  }
</style>
