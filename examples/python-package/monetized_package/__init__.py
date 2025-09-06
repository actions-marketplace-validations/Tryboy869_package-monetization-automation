import re
import requests
from typing import Optional, Dict, Any

class MonetizedPackage:
    def __init__(self, license_key: Optional[str] = None, tier: str = "free"):
        self.license_key = license_key
        self.tier = tier
        self.api_endpoint = "https://api.yourpackage.com"
        
        if not license_key and tier != "free":
            raise ValueError("License key required for paid tiers")
    
    def validate_license(self) -> bool:
        """Validate license key format"""
        if self.tier == "free":
            return True
            
        patterns = {
            "basic": r"^bsc_[a-f0-9]{16}$",
            "pro": r"^pro_[a-f0-9]{16}$", 
            "enterprise": r"^ent_[a-f0-9]{16}$"
        }
        
        pattern = patterns.get(self.tier)
        return bool(pattern and re.match(pattern, self.license_key))
    
    def api_call(self, data: Dict[Any, Any]) -> Dict[Any, Any]:
        """Make API call with usage tracking"""
        if not self.validate_license():
            raise ValueError("Invalid license key")
        
        response = requests.post(
            f"{self.api_endpoint}/process",
            json={
                "data": data,
                "license": self.license_key,
                "tier": self.tier
            }
        )
        
        response.raise_for_status()
        return response.json()