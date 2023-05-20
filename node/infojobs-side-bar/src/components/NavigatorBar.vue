<template>
  <nav class="shadow-lg w-full h-20">
    <div class="max-w-7xl flex justify-between items-center h-full mx-auto">
      <div class="flex gap-8">
        <a class="max-w-sm w-28" href="https://www.infojobs.net" target="_blank" rel="InfoJobs home">
          <img class="" alt="Infojobs logo" src="https://components.infojobs.com/statics/images/infojobs-fill.svg" />
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