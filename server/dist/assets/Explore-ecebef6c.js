import{u as S,b as L,E as b,c as k,r as c,F as o,j as e,L as d,G as v}from"./index-d4ed2424.js";import{u as y,M as C,a as E}from"./MemoResources-3c54ba17.js";import{u as w,t as A}from"./Toast-1e202ccb.js";const D=()=>{const{t:a,i18n:h}=S(),u=L(),x=b(),p=k(),l=y(),[t,i]=c.useState({memos:[]}),[g,n]=c.useState(!1),r=w(),m=u.state.systemStatus.customizedProfile,j=p.state.user,f=x.state;c.useEffect(()=>{l.fetchAllMemos(o,t.memos.length).then(s=>{s.length<o&&n(!0),i({memos:s}),r.setFinish()})},[f]);const N=async()=>{try{const s=await l.fetchAllMemos(o,t.memos.length);s.length<o?n(!0):n(!1),i({memos:t.memos.concat(s)})}catch(s){console.error(s),A.error(s.response.data.message)}};return e.jsx("section",{className:"page-wrapper explore",children:e.jsxs("div",{className:"page-container",children:[e.jsxs("div",{className:"page-header",children:[e.jsxs("div",{className:"title-container",children:[e.jsx("img",{className:"logo-img",src:m.logoUrl,alt:""}),e.jsx("span",{className:"title-text",children:m.name})]}),e.jsx("div",{className:"action-button-container",children:!r.isLoading&&j?e.jsxs(d,{to:"/",className:"link-btn btn-normal",children:[e.jsx("span",{children:"🏠"})," ",a("common.back-to-home")]}):e.jsxs(d,{to:"/auth",className:"link-btn btn-normal",children:[e.jsx("span",{children:"👉"})," ",a("common.sign-in")]})})]}),!r.isLoading&&e.jsxs("main",{className:"memos-wrapper",children:[t.memos.map(s=>{const M=v(s.displayTs).locale(h.language).format("YYYY/MM/DD HH:mm:ss");return e.jsxs("div",{className:"memo-container",children:[e.jsxs("div",{className:"memo-header",children:[e.jsx("span",{className:"time-text",children:M}),e.jsxs("a",{className:"name-text",href:`/u/${s.creator.id}`,children:["@",s.creator.nickname||s.creator.username]})]}),e.jsx(C,{className:"memo-content",content:s.content,onMemoContentClick:()=>{}}),e.jsx(E,{resourceList:s.resourceList})]},s.id)}),g?t.memos.length===0?e.jsx("p",{className:"w-full text-center mt-12 text-gray-600",children:a("message.no-memos")}):null:e.jsx("p",{className:"m-auto text-center mt-4 italic cursor-pointer text-gray-500 hover:text-green-600",onClick:N,children:a("memo-list.fetch-more")})]})]})})};export{D as default};
