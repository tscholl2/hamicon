package main

//generated by embd
const basicCSS = "/* eyes */\n@keyframes blink {\n  50% {\n    ry: 0.1\n  }\n}\n.blink {\n  animation: blink 2s infinite;\n}\n/* mouth */\n@keyframes scaler {\n  30% {\n    transform: scale(0.4)\n  }\n  50% {\n    transform: scale(0.7)\n  }\n  70% {\n    transform: scale(0.5)\n  }\n}\n.talk {\n  animation: scaler 1s infinite alternate;\n}\n#speaker {\n  transform-origin: 50% 50%;\n}\n/* cheeks */\n.swell {\n  animation: scaler 3s infinite;\n}\n#cheek1 {\n  transition: all 0.2s;\n  transform-origin: right center;\n}\n#cheek2 {\n  transition: all 0.2s;\n  transform-origin: left center;\n}\n/* nose */\n@keyframes wiggle {\n  0% {\n    transform: rotate(0deg)\n  }\n  25% {\n    transform: rotate(10deg)\n  }\n  75% {\n    transform: rotate(-10deg)\n  }\n  100% {\n    transform: rotate(0deg)\n  }\n}\n.wiggle {\n  animation: wiggle 1s linear infinite;\n}\n#nose {\n  transform-origin: center bottom;\n}\n/* legs */\n@keyframes walk {\n  0% {\n    transform: rotate(-10deg)\n  }\n  50% {\n    transform: rotate(20deg)\n  }\n  100% {\n    transform: rotate(-10deg)\n  }\n}\n#legs > path {\n  transform-origin: 0% 0%;\n  stroke-linecap: round;\n}\n#legs.walk #bleg1 {\n  animation: walk 1s ease infinite reverse;\n}\n#legs.walk #bleg2 {\n  animation: walk 1s ease 0.35s infinite;\n}\n#legs.walk #fleg1 {\n  animation: walk 1s ease 0.3s infinite reverse;\n}\n#legs.walk #fleg2 {\n  animation: walk 1s ease 0.65s infinite;\n}\n"
