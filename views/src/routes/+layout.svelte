<script>
        import Header from "$lib/Header.svelte"
        import Label from "$lib/Label.svelte"
        import Card from "$lib/Card.svelte"
        import Search from "$lib/Search.svelte"
        import Filter from "$lib/Filter.svelte"
</script>

<Header/>

<span class="label">
     <Search/>
     <Filter/>
     <Label/>   
</span>

<div style="padding: 10px">
        <slot/>
</div>

<style>
        .label {
                display: flex;
                gap: 10px;
                align-items: center;
                overflow: auto;
                padding: 10px;
                margin: 20px 0;
                justify-content: center;
        }
</style>

<svelte:head>
        <meta name="theme-color" content="#000000"/>
        <meta name="apple-mobile-web-app-status-bar-style" content="#000000"/>
        <meta name="msapplication-navbutton-color" content="#000000"/>
</svelte:head>