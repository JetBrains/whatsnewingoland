/**
 * Demo for TypeScript references in JsDoc
 *
 * @type {(name: string, version:string) => IDE}
 */
const go = (name, version) => {
    const i = new IDE();
    i.name = name
    i.version = version
    return i
}

const i = go("GoLand", "2021.2")
console.log(i)
