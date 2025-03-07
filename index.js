import { createIcons, ChevronLeft, ChevronRight} from 'lucide';
import Alpine from 'alpinejs'
import 'htmx.org';

window.Alpine = Alpine
 
Alpine.start()

createIcons({
  icons: {
    ChevronLeft,
    ChevronRight
  }
});