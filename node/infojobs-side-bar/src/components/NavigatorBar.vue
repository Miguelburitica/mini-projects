<template>
  <nav class="shadow-lg w-full h-20">
    <div class="max-w-[1280px] flex justify-between items-center h-full mx-auto">
      <div class="flex gap-8">
        <a class="max-w-sm w-28" href="https://www.infojobs.net" target="_blank" rel="InfoJobs home">
          <svg class="max-w-sm scale-[0.4]" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" x="0" y="0" viewBox="0 0 138 136.054" style="enable-background:new 0 0 138 136.054" xml:space="preserve">
            <path d="M138 112.968c0 12.75-10.335 23.086-23.085 23.086h-91.83C10.336 136.054 0 125.719 0 112.968V23.086C0 10.336 10.336 0 23.085 0h91.83C127.665 0 138 10.336 138 23.086v89.882z" style="fill:#2088c2"></path>
            <path class="st1" d="M57.905 89.323c-.083.749-.748 1.247-1.493 1.247H45.344c-.664 0-1.251-.416-1.251-1.082l.082-.166 7.052-59.253c.085-.666.217-1.004.968-1.087h12.098c.581 0 .914.415.914.997v.166l-7.302 59.178zM86.169 91.298c-1.58 12.898-9.654 15.892-19.39 15.892-5.741 0-6.491-.249-7.988-.415-.748-.085-1.497-.167-1.497-1.332v-.415l.582-4.825c.166-1.415.748-1.662 1.996-1.662 1.165-.083 2.748-.083 4.493-.167 4.325-.248 7.323-.582 8.072-6.743l7.569-61.603c.086-.665.668-1.081 1.416-1.163l11.314.119h.168c.664 0 1.001.415 1.001.998"></path>
          </svg>
        </a>
        <ul class="flex items-center gap-8">
          <li class="cursor-pointer hover:text-blue-400 duration-300">Empleo</li>
          <li class="cursor-pointer hover:text-blue-400 duration-300">Mis ofertas</li>
          <li class="cursor-pointer hover:text-blue-400 duration-300">CV</li>
          <li class="cursor-pointer hover:text-blue-400 duration-300 asp">Quién me ve</li>
        </ul>
      </div>
      <div class="flex items-center gap-3 cursor-pointer">
        <div class="flex justify-center aspect-square w-10 rounded-full overflow-hidden">
          <img class="h-full max-w-none" src="https://images3.alphacoders.com/661/661882.png" alt="profile photo">
        </div>
        <svg class="w-6 h-6 text-gray-300 font-semibold" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={2} stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
        </svg>
      </div>
    </div>
  </nav>
</template>