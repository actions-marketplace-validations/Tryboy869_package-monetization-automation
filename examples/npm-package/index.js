const axios = require('axios');

class MonetizedPackage {
    constructor(options = {}) {
        this.licenseKey = options.licenseKey;
        this.tier = options.tier || 'free';
        this.apiEndpoint = 'https://api.yourpackage.com';
        
        if (!this.licenseKey && this.tier !== 'free') {
            throw new Error('License key required for paid tiers');
        }
    }
    
    async validateLicense() {
        const patterns = {
            'basic': /^bsc_[a-f0-9]{16}$/,
            'pro': /^pro_[a-f0-9]{16}$/,
            'enterprise': /^ent_[a-f0-9]{16}$/
        };
        
        if (this.tier === 'free') return true;
        
        const pattern = patterns[this.tier];
        return pattern && pattern.test(this.licenseKey);
    }
    
    async apiCall(data) {
        if (!await this.validateLicense()) {
            throw new Error('Invalid license key');
        }
        
        // Simulate API call with usage tracking
        const response = await axios.post(`${this.apiEndpoint}/process`, {
            data,
            license: this.licenseKey,
            tier: this.tier
        });
        
        return response.data;
    }
}

module.exports = MonetizedPackage;