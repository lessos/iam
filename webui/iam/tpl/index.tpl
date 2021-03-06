<div class="iam-container" style="padding:10px 0;">
  <div id="iam-topbar">
    <div class="iam-topbar-collapse">
      <ul class="iam-nav" id="iam-topbar-siteinfo">
        <li><img class="iam-topbar-logo" src="/iam/~/iam/img/iam-s2-32.png" title="frontend_header_site_logo_url"></li>
        <li class="iam-topbar-brand">{[=l4i.T(it.pinfo.webui_banner_title)]}</li>
      </ul>
      <ul class="iam-nav iam-topbar-nav" id="iam-topbar-nav-menus">
        <li><a class="l4i-nav-item" href="#index">{[=l4i.T("Home")]}</a></li>
        {[~it.pinfo.topnav :v]}
        <li><a class="l4i-nav-item" href="#{[=v.path]}">{[=v.title]}</a></li>
        {[~]}
      </ul>
      <ul class="iam-nav iam-nav-right" id="iam-topbar-userbar">
        <li><a class="" href="#sign-out" onclick="iamUser.SignOut()">{[=l4i.T("Sign Out")]}</a></li>
      </ul>
    </div>
  </div>
  
  <div id="com-content"></div>

  <div class="iam-footer">
    <img src="/iam/~/iam/img/iam-s2-32.png" class="if-logo"> 
    <a href="https://github.com/hooto/iam" target="_blank">Powered by hooto IAM ({[=l4i.T("Identity & Access Management")]})</a>
  </div>
</div>

