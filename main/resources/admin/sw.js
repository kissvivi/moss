if(!self.define){let s,e={};const l=(l,i)=>(l=new URL(l+".js",i).href,e[l]||new Promise((e=>{if("document"in self){const s=document.createElement("script");s.src=l,s.onload=e,document.head.appendChild(s)}else s=l,importScripts(l),e()})).then((()=>{let s=e[l];if(!s)throw new Error(`Module ${l} didn’t register its module`);return s})));self.define=(i,n)=>{const r=s||("document"in self?document.currentScript.src:"")||location.href;if(e[r])return;let u={};const a=s=>l(s,r),o={module:{uri:r},exports:u,require:a};e[r]=Promise.all(i.map((s=>o[s]||a(s)))).then((s=>(n(...s),u)))}}define(["./workbox-3625d7b0"],(function(s){"use strict";self.skipWaiting(),s.clientsClaim(),s.precacheAndRoute([{url:"assets/@arco-design.55369db5.js",revision:null},{url:"assets/@arco-design.bd8c29bd.css",revision:null},{url:"assets/@codemirror.b45088d6.js",revision:null},{url:"assets/@css-render.6ced7bf3.js",revision:null},{url:"assets/@emotion.6322e2ae.js",revision:null},{url:"assets/@intlify.bed9fa1a.js",revision:null},{url:"assets/@juggle.32c34d6c.js",revision:null},{url:"assets/@lezer.133a6ecb.js",revision:null},{url:"assets/@vue.0987707a.js",revision:null},{url:"assets/@vueuse.d5398ce4.js",revision:null},{url:"assets/@wangeditor.501cf061.css",revision:null},{url:"assets/@wangeditor.ba1b7edd.js",revision:null},{url:"assets/admin.650d9c8f.js",revision:null},{url:"assets/api.17b3db47.js",revision:null},{url:"assets/async-validator.66a9c2c1.js",revision:null},{url:"assets/axios.b9f958b0.js",revision:null},{url:"assets/b-tween.87ffe365.js",revision:null},{url:"assets/b-validate.ee581f7d.js",revision:null},{url:"assets/b2.3ae607bd.js",revision:null},{url:"assets/badger.a4060a3a.js",revision:null},{url:"assets/cache.84762293.js",revision:null},{url:"assets/codemirror.7821c22f.js",revision:null},{url:"assets/compute-scroll-into-view.17358474.js",revision:null},{url:"assets/cos.3d03095f.js",revision:null},{url:"assets/createAdmin.e253ff6e.js",revision:null},{url:"assets/crelt.67277586.js",revision:null},{url:"assets/css-render.20ab466e.js",revision:null},{url:"assets/date-fns-tz.c3c7eb03.js",revision:null},{url:"assets/date-fns.a06005bd.js",revision:null},{url:"assets/dayjs.396bdce9.js",revision:null},{url:"assets/DetectLinks.046cf767.css",revision:null},{url:"assets/DetectLinks.130160d5.js",revision:null},{url:"assets/Duration.f6f560b6.js",revision:null},{url:"assets/evtd.9eee5233.js",revision:null},{url:"assets/ftp.fe153d39.js",revision:null},{url:"assets/GenerateDescription.bc79cec6.js",revision:null},{url:"assets/GenerateSlug.2cec3680.js",revision:null},{url:"assets/index.052bb61f.css",revision:null},{url:"assets/index.0a05e831.js",revision:null},{url:"assets/index.17074d49.js",revision:null},{url:"assets/index.1708b6bf.js",revision:null},{url:"assets/index.1a1de342.js",revision:null},{url:"assets/index.2e19c31e.css",revision:null},{url:"assets/index.33e324a7.css",revision:null},{url:"assets/index.39edbefa.js",revision:null},{url:"assets/index.58a1dc4e.js",revision:null},{url:"assets/index.77599200.js",revision:null},{url:"assets/index.79277f34.css",revision:null},{url:"assets/index.7f1e5104.css",revision:null},{url:"assets/index.8b92274b.js",revision:null},{url:"assets/index.8ebd22e4.css",revision:null},{url:"assets/index.ac3f92ff.js",revision:null},{url:"assets/index.b3e7e476.js",revision:null},{url:"assets/index.b5c10ed9.css",revision:null},{url:"assets/index.ee4a9256.css",revision:null},{url:"assets/layout.99a3c5db.css",revision:null},{url:"assets/layout.bbbe4b94.css",revision:null},{url:"assets/layout.e8c415c2.js",revision:null},{url:"assets/layout.f810aeb5.js",revision:null},{url:"assets/local.c148f298.js",revision:null},{url:"assets/lodash-es.33d1f95f.js",revision:null},{url:"assets/log.e9967ebc.js",revision:null},{url:"assets/login.64683044.js",revision:null},{url:"assets/memcached.78b94051.js",revision:null},{url:"assets/more.3ee8fb03.js",revision:null},{url:"assets/naive-ui.506a45a0.js",revision:null},{url:"assets/number-precision.6dad9ff9.js",revision:null},{url:"assets/oss.18f9e363.js",revision:null},{url:"assets/pinia.2e07300c.js",revision:null},{url:"assets/prettier.5505bad8.js",revision:null},{url:"assets/redis.f1b9700c.js",revision:null},{url:"assets/resize-observer-polyfill.8deb1e21.js",revision:null},{url:"assets/router.13dace6b.js",revision:null},{url:"assets/s3.63170420.js",revision:null},{url:"assets/SaveArticleImages.45e6e1fa.css",revision:null},{url:"assets/SaveArticleImages.b37df7eb.js",revision:null},{url:"assets/scroll-into-view-if-needed.61c672a4.js",revision:null},{url:"assets/seemly.d0f7d7a4.js",revision:null},{url:"assets/site.4755efb5.js",revision:null},{url:"assets/sitemap.6b677654.js",revision:null},{url:"assets/source-map.205bdfab.js",revision:null},{url:"assets/style-mod.a2e40363.js",revision:null},{url:"assets/subMenu.0d5c58da.js",revision:null},{url:"assets/subMenu.10a3917e.css",revision:null},{url:"assets/tailwindcss.4305d109.css",revision:null},{url:"assets/template.899bc9e4.js",revision:null},{url:"assets/template.8e838242.css",revision:null},{url:"assets/theme.6bbfd5ed.css",revision:null},{url:"assets/theme.fb3c3ce1.js",revision:null},{url:"assets/treemate.ca338995.js",revision:null},{url:"assets/upload.b741c949.js",revision:null},{url:"assets/UploadImgInput.ea0b5d48.js",revision:null},{url:"assets/url.1d176357.js",revision:null},{url:"assets/url.c246d41b.css",revision:null},{url:"assets/vdirs.ab69c576.js",revision:null},{url:"assets/vooks.3f61458b.js",revision:null},{url:"assets/vue-codemirror.ecdbe3ea.js",revision:null},{url:"assets/vue-demi.b3a9cad9.js",revision:null},{url:"assets/vue-i18n.e3137642.js",revision:null},{url:"assets/vue-request.6886b8d7.js",revision:null},{url:"assets/vue-router.0ed66d6f.js",revision:null},{url:"assets/vue.5c5bb0aa.js",revision:null},{url:"assets/vueuc.5f5811a3.js",revision:null},{url:"assets/w3c-keyname.30cf5eb3.js",revision:null},{url:"index.html",revision:"7624efb9635aad4fa2f6f21b903953df"},{url:"loading.css",revision:"42742bd6ee5361d2dbd830c4d3c17464"},{url:"registerSW.js",revision:"b878d4c5a304cf9402789989e4ad58f2"},{url:"favicon.ico",revision:"2ee8f9467e1bafebb07a73f19cfe4ce3"},{url:"apple-touch-icon.png",revision:"687a5e567e381abdeff7378b4a27c87f"},{url:"mask_icon.svg",revision:"34f94be4aa9de3aef5b19a35c5855ee7"},{url:"icon_192.png",revision:"687a5e567e381abdeff7378b4a27c87f"},{url:"icon_512.png",revision:"34957835788f7109e4ce98eb3d5b6914"},{url:"manifest.webmanifest",revision:"b74d9be4738bc285e9758c4237611609"}],{}),s.cleanupOutdatedCaches(),s.registerRoute(new s.NavigationRoute(s.createHandlerBoundToURL("index.html")))}));
