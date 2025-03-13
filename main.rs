use rand::Rng;
use std::f64::consts::E;

struct Colony {
    trust: f64,
    trust_min: f64,
    success_rate: f64,
}

const ALPHA: f64 = 0.1; // Learning rate for trust growth
const BETA: f64 = 0.05; // Decay factor (for alternative version)
const GAMMA: f64 = 0.02; // Adaptive threshold scaling factor
const MAX_ITERATIONS: usize = 100;

impl Colony {
    fn update_trust(&mut self, success: bool) {
        if success {
            self.success_rate += 1.0;
            self.trust += (1.0 - self.trust) * (1.0 - E.powf(-ALPHA * self.success_rate));
        } else {
            self.trust *= E.powf(-BETA);
        }
        
        // Adaptive minimum trust
        self.trust_min = (0.1 + GAMMA * self.success_rate).min(1.0);
        self.trust = self.trust.max(self.trust_min);
    }
}

fn main() {
    let mut rng = rand::thread_rng();
    let mut colony = Colony {
        trust: 0.1,
        trust_min: 0.1,
        success_rate: 0.0,
    };

    for i in 1..=MAX_ITERATIONS {
        let ka_success = rng.gen::<f64>() < 0.7; // Simulate KA's success rate
        colony.update_trust(ka_success);
        println!("Iteration {}: Trust = {:.4}, TrustMin = {:.4}", i, colony.trust, colony.trust_min);
    }
}