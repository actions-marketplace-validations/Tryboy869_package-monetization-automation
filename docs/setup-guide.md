# Setup Guide

## Quick Start

1. Add the action to your workflow:

```yaml
- uses: Tryboy869/package-monetization-automation@v1
  with:
    monetization_type: 'usage_based'
    pricing_tier: 'basic'
```

2. Configure payment processor (Stripe recommended)
3. Set up secrets for API keys
4. Test with sample integration

## Detailed Configuration

### Payment Processors

#### Stripe Setup
1. Create Stripe account
2. Get API keys from dashboard
3. Add to GitHub secrets as `STRIPE_API_KEY`

#### PayPal Setup
1. Create PayPal developer account
2. Create application
3. Add credentials to secrets

### Pricing Configuration

You can customize pricing with:

```yaml
- uses: Tryboy869/package-monetization-automation@v1
  with:
    custom_pricing: |
      {
        "free": {"price": 0, "limit": 1000},
        "basic": {"price": 25, "limit": 15000}
      }
```
