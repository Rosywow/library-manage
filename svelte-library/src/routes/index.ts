import { wrap } from "svelte-spa-router/wrap"


const modules: any = import.meta.glob('@/modules/**/*.route.ts', {
    eager: true,
});


let modulesObject = {};
for (const path in modules) {
    console.log(modules[path].default)
    modulesObject = { ...modulesObject, ...modules[path].default }
}
// console.log("Array:", modulesObject)

const router = {
    ...modulesObject
}
export default router;