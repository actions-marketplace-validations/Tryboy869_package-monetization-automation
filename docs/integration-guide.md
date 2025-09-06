# Integration Guide

## For Package Users

### NPM Integration
```javascript
const MyPackage = require('your-package');
const client = new MyPackage({
  licenseKey: 'pro_1234567890abcdef',
  tier: 'pro'
});
```

### Python Integration  
```python
from your_package import MonetizedPackage
client = MonetizedPackage(
    license_key='pro_1234567890abcdef',
    tier='pro'
)
```

### Go Integration
```go
import "github.com/you/your-package"

client, err := yourpackage.New("pro_1234567890abcdef", "pro")
```

## For Package Developers

### Adding Usage Tracking
```javascript
// Track API usage (privacy-compliant)
function trackUsage(endpoint, tier) {
    // Implementation depends on your analytics choice
    console.log(`API call: ${endpoint}, tier: ${tier}`);
}
```

### License Validation
All examples include license validation patterns for each tier.
