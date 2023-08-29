(function(){"use strict";self.importScripts("./wasm_exec.js"),WebAssembly.instantiateStreaming||(WebAssembly.instantiateStreaming=async(e,t)=>{const a=await(await e).arrayBuffer();return await WebAssembly.instantiate(a,t)});function r(e){const t=new Go;WebAssembly.instantiateStreaming(fetch(e.wasm),t.importObject).then(a=>{t.run(a.instance),postMessage({type:n.Ready})}).catch(a=>{console.error(a),postMessage({type:n.Failed,reason:a instanceof Error?a.message:"Unknown Error"})})}function s(e){const t=initializeWorker(e.cfg);return t!=null?{type:n.Failed,reason:JSON.parse(t).error}:{type:n.Initialized}}function o(e){const t=simulate();return typeof t=="string"||t instanceof String?{type:n.Failed,reason:JSON.parse(t).error}:{type:n.Done,result:t,itr:e.itr}}function u(e){switch(e.type){case i.Ready:return r(e);case i.Initialize:return postMessage(s(e));case i.Run:return postMessage(o(e));default:throw console.error("aggregator - unknown request: ",e),new Error("aggregator unknown request")}}onmessage=e=>u(e.data);var i=(e=>(e.Ready="ready",e.Initialize="initialize",e.Run="run",e))(i||{}),n=(e=>(e.Failed="failed",e.Ready="ready",e.Initialized="initialized",e.Done="done",e))(n||{})})();
