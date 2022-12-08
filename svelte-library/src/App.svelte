<script>
    // 主要处理的相关绑定
    import Router from "svelte-spa-router";
    import routes from "./routes/index";
</script>

<main>
    <Router {routes} />
</main>
